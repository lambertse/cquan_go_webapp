package seeds

import (
	"log"

	"github.com/lambertse/cquan_go_webapp/internal/models"
	"gorm.io/gorm"
)

func SeedUser(DB *gorm.DB) error {
  	// Create default roles if they don't exist
	roles := []models.Role{
		{Name: "admin", Description: "Administrator with full access"},
		{Name: "user", Description: "Standard user with basic access"},
	}

	for _, role := range roles {
		var existingRole models.Role
		result := DB.Where("name = ?", role.Name).First(&existingRole)
		if result.Error != nil {
			if result.Error.Error() == "record not found" {
				if err := DB.Create(&role).Error; err != nil {
					return err
				}
				log.Printf("Created role: %s", role.Name)
			} else {
				return result.Error
			}
		}
	}


	// Assign all permissions to admin role
	var adminRole models.Role
	if err := DB.Where("name = ?", "admin").First(&adminRole).Error; err != nil {
		return err
	}

	// Create default admin user if it doesn't exist
	var adminUser models.User
	result := DB.Where("username = ?", "admin").First(&adminUser)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			adminUser = models.User{
				Username: "admin",
				Password: "123", // This will be hashed by the BeforeSave hook
			}
			if err := DB.Create(&adminUser).Error; err != nil {
				return err
			}
			log.Printf("Created admin user: %s", adminUser.Username)

			// Assign admin role to admin user
			if err := DB.Model(&adminUser).Association("Roles").Append(&adminRole); err != nil {
				return err
			}
			log.Printf("Assigned admin role to admin user")
		} else {
			return result.Error
		}
	}

  // Crate default user if it doesn't exist
  var defaultUser models.User
  result = DB.Where("username = ?", "user").First(&defaultUser)
  if result.Error != nil {
    if result.Error.Error() == "record not found" {
      defaultUser = models.User{
        Username: "user",
        Password: "123", // This will be hashed by the BeforeSave hook
      }
      if err := DB.Create(&defaultUser).Error; err != nil {
        return err
      }
      log.Printf("Created default user: %s", defaultUser.Username)
      // Assign user role to default user 
      var userRole models.Role
      if err := DB.Where("name = ?", "user").First(&userRole).Error; err != nil {
        return err
      }
      if err := DB.Model(&defaultUser).Association("Roles").Append(&userRole); err != nil {
        return err
      }
      log.Printf("Assigned user role to default user")
    } else {
      return result.Error
    }

  }

	return nil
}
