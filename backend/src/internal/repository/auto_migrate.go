package repository

import (
	"llmapi/src/internal/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	// Automatically migrate the schema, creating tables, constraints, columns, and indexes as needed.
	// This is a safe operation, but will change existing columns if they are not compatible with the new schema.
	if err := db.AutoMigrate(
		&model.User{},
		&model.APIKeyRecord{},
	); err != nil {
		return err
	}
	return nil
}
