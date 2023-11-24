package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/adindazenn/kelompok6/model"
	"github.com/adindazenn/kelompok6/database"
)

// Endpoint untuk membuat Task (POST /tasks)
func CreateTask(c *gin.Context) {
    // Ambil user dari token JWT
    user, err := GetUserFromToken(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
        return
    }

    // Bind request body ke struct Task
    var task model.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Cek apakah kategori dengan ID yang diberikan ada dalam basis data
    db, err := database.InitDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengakses database"})
        return
    }

    var category model.Category
    if err := db.Where("id = ?", task.CategoryID).First(&category).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Kategori tidak ditemukan"})
        return
    }

    // Set nilai-nilai yang diperlukan
    task.UserID = user.ID
    task.Status = false
    task.CreatedAt = time.Now()

    // Simpan task ke database
    if err := db.Create(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat task"})
        return
    }

    Response := model.CreateTaskResponse{
        ID: task.ID,       
        Title:  task.Title,     
        Status: task.Status,     
        Description: task.Description,
        UserID:     task.UserID, 
        CategoryID: task.CategoryID,
        CreatedAt:  task.CreatedAt,
    }
    c.JSON(http.StatusCreated, Response)
}

func getUserByID(userID uint) (model.User, error) {
    db, err := database.InitDB()
    if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengakses database"})
	return
    }
    var user model.User
    if userErr := db.Where("id = ?", userID).First(&user).Error; err != nil {
        return model.User{}, userErr
    }
    return user, nil
}

func GetTasks(c *gin.Context) {
    // Ambil user dari token JWT
    _, err := GetUserFromToken(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
        return
    }

    // Ambil daftar tasks dari database
    db, err := database.InitDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengakses database"})
        return
    }

    var tasks []model.Task
    if err := db.Find(&tasks).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data tugas"})
        return
    }

    // Membuat respons menggunakan struct TaskResponse
    var response []model.TaskResponse
    for _, task := range tasks {
	    dataUser, dataErr := getUserByID(task.UserID)
	        if dataErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data user"})
        	return
	    }
        taskResponse := model.TaskResponse{
            ID:          task.ID,
            Title:       task.Title,
            Status:      task.Status,
            Description: task.Description,
            UserID:      task.UserID,
            CategoryID:  task.CategoryID,
            CreatedAt:   task.CreatedAt,
        }
        taskResponse.User.ID = dataUser.ID
        taskResponse.User.Email = dataUser.Email
        taskResponse.User.FullName = dataUser.FullName
        response = append(response, taskResponse)
    }

    c.JSON(http.StatusOK, response)
}

func UpdateTask(c *gin.Context) {
    // Ambil user dari token JWT
    user, err := GetUserFromToken(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
        return
    }

    // Ambil task ID dari parameter
    taskID := c.Param("taskId")

    // Ambil tugas dari database
    db, err := database.InitDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengakses database"})
        return
    }
    var task model.Task
    if err := db.Where("id = ? AND user_id = ?", taskID, user.ID).First(&task).Error; err != nil {
        c.JSON(http.StatusForbidden, gin.H{"error": "Anda tidak diizinkan melakukan tindakan ini"})
        return
    }

    // Bind data permintaan ke tugas yang ada
    var request model.UpdateTaskRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update data tugas
    task.Title = request.Title
    task.Description = request.Description
    task.UpdatedAt = time.Now()

    if err := db.Save(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui tugas"})
        return
    }

    response := model.TaskUpdateResponse{
        ID:          task.ID,
        Title:       task.Title,
        Description: task.Description,
        Status:      task.Status,
        UserID:      task.UserID,
        CategoryID:  task.CategoryID,
        UpdatedAt:   task.UpdatedAt,
    }

    c.JSON(http.StatusOK, response)
}

// Endpoint untuk memperbarui status task (PATCH /tasks/update-status/:taskId)
func UpdateTaskStatus(c *gin.Context) {
    // Ambil user dari token JWT
    user, err := GetUserFromToken(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
        return
    }

    // Ambil taskId dari parameter URL
    taskId := c.Param("taskId")

    // Cari task dalam database
    db, err := database.InitDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengakses database"})
        return
    }

    var task model.Task
    if err := db.Where("id = ? AND user_id = ?", taskId, user.ID).First(&task).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task tidak ditemukan"})
        return
    }

    // Bind request body ke struct UpdateTaskStatusRequest
    var request model.UpdateTaskStatusRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update status task
    task.Status = request.Status

    // Simpan perubahan ke database
    if err := db.Save(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui status task"})
        return
    }

    response := model.TaskUpdateResponse{
        ID:          task.ID,
        Title:       task.Title,
        Description: task.Description,
        Status:      task.Status,
        UserID:      task.UserID,
        CategoryID:  task.CategoryID,
        UpdatedAt:   task.UpdatedAt,
    }

    c.JSON(http.StatusOK, response)
}

// Endpoint untuk memperbarui category task (PATCH /tasks/update-category/:taskId)
func UpdateTaskCategory(c *gin.Context) {
    // Ambil user dari token JWT
    user, err := GetUserFromToken(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
        return
    }

    // Ambil taskId dari parameter URL
    taskId := c.Param("taskId")

    // Cari task dalam database
    db, err := database.InitDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengakses database"})
        return
    }

    var task model.Task
    if err := db.Where("id = ? AND user_id = ?", taskId, user.ID).First(&task).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task tidak ditemukan"})
        return
    }

    // Bind request body ke struct UpdateTaskCategoryRequest
    var request model.UpdateTaskCategoryRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Periksa apakah category dengan ID yang diberikan oleh pengguna ada di database
    var category model.Category
    if err := db.Where("id = ?", request.CategoryID).First(&category).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Category tidak ditemukan"})
        return
    }

    // Update category task
    task.CategoryID = request.CategoryID

    // Simpan perubahan ke database
    if err := db.Save(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui category task"})
        return
    }

    response := model.TaskUpdateResponse{
        ID:          task.ID,
        Title:       task.Title,
        Description: task.Description,
        Status:      task.Status,
        UserID:      task.UserID,
        CategoryID:  task.CategoryID,
        UpdatedAt:   task.UpdatedAt,
    }

    c.JSON(http.StatusOK, response)
}

// DeleteTask adalah endpoint untuk menghapus task milik pengguna (DELETE /tasks/:taskId)
func DeleteTask(c *gin.Context) {
    // Ambil user dari token JWT
    user, err := GetUserFromToken(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
        return
    }

    // Ambil taskId dari parameter URL
    taskID := c.Param("taskId")

    db, err := database.InitDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengakses database"})
        return
    }

    var task model.Task
    if err := db.First(&task, taskID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task tidak ditemukan"})
        return
    }

    // Cek apakah pengguna yang menghapus task adalah pemiliknya
    if task.UserID != user.ID {
        c.JSON(http.StatusForbidden, gin.H{"error": "Anda tidak diizinkan menghapus task ini"})
        return
    }

    // Hapus task dari database
    if err := db.Delete(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus task"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task has been succesfully deleted"})
}



