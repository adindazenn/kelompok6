// main.go

package main

import (
    "github.com/adindazenn/kelompok6/router" 
    "github.com/adindazenn/kelompok6/database"
	"github.com/adindazenn/kelompok6/model"
	"fmt"
	"os"
    "golang.org/x/crypto/bcrypt"
    "log"
)

func seedAdmin() error {
	// Inisialisasi koneksi ke database
    db, err := database.InitDB()
    if err != nil {
        fmt.Println("Error initializing database:", err)
        log.Fatal(err)
    }

    // Cek apakah ada admin di database
    var admin model.User
    if db.Where("role = ?", "admin").First(&admin); admin.ID == 0 {
        // Jika tidak ada admin, buat satu admin baru
        hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
        newAdmin := model.User{
            FullName: "Admin",
            Email:    "admin@example.com",
            Password: string(hashedPassword),
            Role:     "admin",
        }
        if err := db.Create(&newAdmin).Error; err != nil {
            log.Fatal("Error creating admin:", err)
        }
        fmt.Println("Admin seeded successfully")
    } else if db.Error != nil {
        log.Fatal("Error querying admin:", db.Error)
    }

	return nil
}

func main() {
	var PORT = os.Getenv("PORT")
    r := router.SetupRouter()
    // Jalankan seeding data admin
	if err := seedAdmin(); err != nil {
		fmt.Println("Error seeding admin data:", err)
		return
	}
    r.Run(":" + PORT) 
}
