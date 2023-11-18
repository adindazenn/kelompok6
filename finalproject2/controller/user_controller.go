// controller/user_controller.go

package controller

import (
    "net/http"
    "time"
	"fmt"
    "errors"

    "github.com/gin-gonic/gin"
    "github.com/adindazenn/kelompok6/finalproject2/model"  
	"github.com/adindazenn/kelompok6/finalproject2/database"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
)

func RegisterUser(c *gin.Context) {
    db, err := database.InitDB()
    if err != nil {
        fmt.Println("Error initializing database:", err)
        return
    }

    var newUser model.User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newUser.Role = "member"
    newUser.CreatedAt = time.Now()

    // Hash kata sandi sebelum menyimpannya di database
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal meng-hash kata sandi"})
        return
    }
    newUser.Password = string(hashedPassword)

    // Simpan user baru ke database
    if err := db.Create(&newUser).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Response data user yang telah terdaftar
    userResponse := model.UserResponse{
        ID:        newUser.ID,
        FullName:  newUser.FullName,
        Email:     newUser.Email,
        CreatedAt: newUser.CreatedAt,
    }
    c.JSON(http.StatusCreated, userResponse)
}

func LoginUser(c *gin.Context) {
    db, err := database.InitDB()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
    var request model.LoginRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Search user berdasarkan alamat email
    var user model.User
    if err := db.Where("email = ?", request.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Email tidak ditemukan"})
        return
    }

    // Verifikasi kata sandi
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Kata sandi salah"})
        return
    }

    // Buat token JWT
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["user_id"] = user.ID
    tokenString, err := token.SignedString([]byte("my_secret_key")) // Ganti dengan kunci rahasia Anda

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
        return
    }

    // Response dengan token JWT
    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func GetUserFromToken(c *gin.Context) (*model.User, error) {
    db, err := database.InitDB()
    if err != nil {
        fmt.Println("Error initializing database:", err)
        return nil, err
    }
    
    // Mendapatkan token dari header Authorization
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        return nil, errors.New("Token JWT tidak ditemukan")
    }

    // Menghapus "Bearer " dari token
    tokenString := authHeader[len("Bearer "):]

    // Memeriksa dan memverifikasi token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte("my_secret_key"), nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        // Mendapatkan user_id dari klaim token
        user_id, ok := claims["user_id"].(float64)
        if !ok {
            return nil, errors.New("User ID tidak valid")
        }

        // Temukan pengguna berdasarkan user_id
        var user model.User
        if err := db.Where("id = ?", int(user_id)).First(&user).Error; err != nil {
            return nil, err
        }

        return &user, nil
    } else {
        return nil, errors.New("Token JWT tidak valid")
    }
}


func UpdateAccount(c *gin.Context) {
    db, err := database.InitDB()
    if err != nil {
        fmt.Println("Error initializing database:", err)
        return 
    }
    // Get user yang terotentikasi dari token JWT
    user, err := GetUserFromToken(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
        return
    }

    // Parse data permintaan update akun
    var request model.UpdateAccountRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update informasi akun user
    user.FullName = request.FullName
    user.Email = request.Email

    // Update data user dalam database
    if err := db.Save(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui akun"})
        return
    }

    // Response data user yang telah terdaftar
    userResponse := model.UpdateUserResponse{
        ID:        user.ID,
        FullName:  user.FullName,
        Email:     user.Email,
        UpdatedAt: user.UpdatedAt,
    }
    c.JSON(http.StatusCreated, userResponse)
}

func DeleteAccount(c *gin.Context) {
    db, err := database.InitDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengakses database"})
        return
    }

    // get user yang terotentikasi dari token JWT
    user, err := GetUserFromToken(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Autentikasi gagal"})
        return
    }

    // Hapus akun user
    if err := db.Delete(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus akun"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Your account has been successfully deleted"})
}

