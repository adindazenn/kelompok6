package controller

import (
	"net/http"
	"time"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/adindazenn/kelompok6/model"
	"github.com/adindazenn/kelompok6/database"
)

// Endpoint untuk membuat Kategori (POST /categories)
func CreateCategory(c *gin.Context) {
	// Ambil user dari token JWT
	user, err := GetUserFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
		return
	}

	// Cek apakah user memiliki role admin
	if user.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Anda tidak diizinkan melakukan tindakan ini"})
		return
	}

	// Bind request body ke struct Category
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set waktu pembuatan
	category.CreatedAt = time.Now()

	// Simpan kategori ke database
	db, err := database.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengakses database"})
		return
	}
	if err := db.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat kategori"})
		return
	}

    Response := model.CreateCategoryResponse{
        ID:        category.ID,
        Type:     category.Type,
        CreatedAt: category.CreatedAt,
    }
    c.JSON(http.StatusCreated, Response)
}

// Endpoint untuk mendapatkan daftar Kategori (GET /categories)
func GetCategories(c *gin.Context) {
	// Ambil user dari token JWT
	_, err := GetUserFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
		return
	}

	// Ambil daftar kategori dari database
	db, err := database.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengakses database"})
		return
	}

	var categories []model.Category
	if err := db.Preload("Tasks").Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data kategori"})
		return
	}

	// Buat struktur CategoryResponse dengan data yang sesuai
	categoryResponses := make([]model.CategoryResponse, len(categories))
	for i, category := range categories {
		tasks := []model.TaskInfo{} // Isi dengan data tugas yang sesuai
		categoryResponses[i] = model.CategoryResponse{
			ID:        category.ID,
			Type:      category.Type,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
			Tasks:     tasks,
		}
	}

	c.JSON(http.StatusOK, categoryResponses)
}

// Endpoint untuk mengupdate Kategori (PATCH /categories/:categoryId)
func UpdateCategory(c *gin.Context) {
	// Ambil user dari token JWT
	user, err := GetUserFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
		return
	}

	// Cek apakah user memiliki role admin
	if user.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Anda tidak diizinkan melakukan tindakan ini"})
		return
	}

	// Ambil ID kategori dari parameter
	categoryID := c.Param("categoryId")

	// Bind request body ke struct Category
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan waktu update
	category.UpdatedAt = time.Now()

	// Update kategori dalam database
	db, err := database.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengakses database"})
		return
	}

	if err := db.Model(&model.Category{}).Where("id = ?", categoryID).Updates(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui kategori"})
		return
	}

	ID, err := strconv.ParseUint(categoryID, 10, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengconvert ID"})
		return
	}

    Response := model.UpdateCategoryResponse{
        ID: uint(ID),
        Type:     category.Type,
        UpdatedAt: category.UpdatedAt,
    }
    c.JSON(http.StatusCreated, Response)
}

// Endpoint untuk menghapus Kategori (DELETE /categories/:categoryId)
func DeleteCategory(c *gin.Context) {
    // Ambil user dari token JWT
    user, err := GetUserFromToken(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
        return
    }

    // Cek apakah user memiliki role admin
    if user.Role != "admin" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Anda tidak diizinkan melakukan tindakan ini"})
        return
    }

    // Ambil categoryId dari parameter URL
    categoryId := c.Param("categoryId")

    // Hapus kategori dari database berdasarkan categoryId
    db, err := database.InitDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengakses database"})
        return
    }

    if err := db.Where("id = ?", categoryId).Delete(&model.Category{}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus kategori"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Category has been successfully deleted"})
}


