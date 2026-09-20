package seeder

import (
	"errors"
	"log"

	"github.com/mohamed-karam/go-food-delivery/identity-service/entity"
	"github.com/mohamed-karam/go-food-delivery/identity-service/pkg"
	"gorm.io/gorm"
)


func SeedUsers(db *gorm.DB) error {
    var adminRole entity.Role
    var customerRole entity.Role
    var driverRole entity.Role

    if err := db.Where("name = ?", "Admin").First(&adminRole).Error; err != nil {
        return err
    }

    if err := db.Where("name = ?", "Customer").First(&customerRole).Error; err != nil {
        return err
    }

    if err := db.Where("name = ?", "Driver").First(&driverRole).Error; err != nil {
        return err
    }

    users := []entity.User{
        {
            Name:           "Admin",
            Email:          "admin@gmail.com",
            Phone:          "01202095030",
            PasswordHash:   hashPassword("password"),
            RoleID:         adminRole.ID,
            Status:         "Active",
        },
        {
            Name:           "Customer",
            Email:          "customer@gmail.com",
            Phone:          "01211095030",
            PasswordHash:   hashPassword("password"),
            RoleID:         customerRole.ID,
            Status:         "Active",
        },
        {
            Name:           "Driver",
            Email:          "driver@gmail.com",
            Phone:          "01211093330",
            PasswordHash:   hashPassword("password"),
            RoleID:         driverRole.ID,
            Status:         "Active",
        },
    }

    log.Printf(
        "Roles found: Admin=%d, Customer=%d, Driver=%d",
        adminRole.ID,
        customerRole.ID,
        driverRole.ID,
    )

    for _, user := range users {
        var existingUser entity.User
        err := db.Where("email = ?", user.Email).First(&existingUser).Error
        if err == nil {
			log.Printf("⏭️ User already exists: %s", user.Email)
            continue
        }

        if !errors.Is(err, gorm.ErrRecordNotFound) {
            return err
        }

        log.Printf("Creating user: %s with RoleID=%d", user.Email, user.RoleID)
        if err := db.Create(&user).Error; err != nil {
            return err
        }

        log.Printf("✅ User seeded: %s", user.Email)
    }

    return nil
}


func hashPassword(password string) string {
    hashedPassword, err := pkg.HashPassword(password)
    if err != nil {
        panic(err)
    }
    return string(hashedPassword)
}