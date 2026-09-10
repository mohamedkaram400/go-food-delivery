package seeder

import (
	"log"

	"github.com/mohamed-karam/go-food-delivery/identity-service/entity"
	"gorm.io/gorm"
)


func SeedRoles(db *gorm.DB) error {

    roles := []entity.Role{
        {
            ID:             1,
            Name:           "Admin",
        },
        {
            ID:             2,
            Name:           "Customer",
        },
        {
            ID:             3,
            Name:           "Driver",
        },
    }

    for _, role := range roles {
        var existingRole entity.Role
        err := db.Where("name = ?", role.Name).First(&existingRole).Error
        if err != nil {
			log.Printf("⏭️ Role already exists: %s", role.Name)
            continue
        }

        if err != gorm.ErrRecordNotFound {
            return err
        }

        if err := db.Create(&role).Error; err != nil {
            return err
        }

        log.Printf("✅ Role seeded: %s", role.Name)
    }

    return nil
}

