package db

import (
	"github.com/lambertse/cquan_go_webapp/internal/db/seed"
	"github.com/lambertse/cquan_go_webapp/internal/models"
)

// MigrateAndSeed performs database migration and seeds initial data
func MigrateAndSeed() error {
	// Auto migrate models
	err := DB.AutoMigrate(
		&models.User{},
		&models.Role{},
	)
	if err != nil {
		return err
	}

	// Seed default data
	err = seedDefaultData()
	if err != nil {
		return err
	}

	return nil
}

// seedDefaultData seeds the database with default data
func seedDefaultData() error {
  err := seeds.SeedUser(DB)
  if err != nil {
    return err
  }

  return nil
}
