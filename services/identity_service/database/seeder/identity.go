package seeder

import (
	"log"

	"github.com/mohamed-karam/go-food-delivery/identity-service/entity"
	"github.com/mohamed-karam/go-food-delivery/identity-service/pkg"
	"gorm.io/gorm"
)


func SeedUsers(db *gorm.DB) error {

    users := []entity.User{
        {
            Name:           "Admin",
            Email:          "admin@gmail.com",
            Phone:          stringPtr("01202095030"),
            PasswordHash:   hashPassword("password"),
            RoleID:         1,
            Status:         "Active",
        },
        {
            Name:           "Admin",
            Email:          "admin@gmail.com",
            Phone:          stringPtr("01202095030"),
            PasswordHash:   hashPassword("password"),
            RoleID:         2,
            Status:         "Active",
        },
    }

    for _, user := range users {
        var existingUser entity.User
        err := db.Where("email = ?", user.Email).First(&existingUser).Error
        if err == nil {
			log.Printf("⏭️ User already exists: %s", user.Email)
            continue
        }

        if err != gorm.ErrRecordNotFound {
            return err
        }

        if err := db.Create(&user).Error; err != nil {
            return err
        }

        log.Printf("✅ User seeded: %s", user.Email)
    }

    return nil
}


func stringPtr(value string) *string {
	return &value
}

func hashPassword(password string) string {
    hashedPassword, err := pkg.HashPassword(password)
    if err != nil {
        panic(err)
    }
    return string(hashedPassword)
}