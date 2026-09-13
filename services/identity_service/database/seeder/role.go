package seeder

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/mohamed-karam/go-food-delivery/identity-service/entity"
	"gorm.io/gorm"
)


func SeedRoles(db *gorm.DB) error {

    roles := []entity.Role{
        {
            ID:             1,
            Name:           "Admin",
            CreatedAt:      time.Now(),
        },
        {
            ID:             2,
            Name:           "Customer",
            CreatedAt:      time.Now(),
        },
        {
            ID:             3,
            Name:           "Driver",
            CreatedAt:      time.Now(),
        },
    }

    for _, role := range roles {
        var existingRole entity.Role
        err := db.Where("name = ?", role.Name).First(&existingRole).Error
        if err == nil {
			log.Printf("⏭️ Role already exists: %s", role.Name)
            continue
        }

        fmt.Println(err)
        if !errors.Is(err, gorm.ErrRecordNotFound) {
            return err
        }

        log.Printf("Creating role: %s", role.Name)

        if err := db.Create(&role).Error; err != nil {
            return err
        }

        log.Printf("✅ Role seeded: %s", role.Name)
    }

    return nil
}

