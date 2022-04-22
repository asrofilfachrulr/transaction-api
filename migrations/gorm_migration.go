package migrations

import (
	"os"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if os.Getenv("DROP") == "y" {
		db.Migrator().DropTable(Models...)
	}

	if err := db.AutoMigrate(Models...); err != nil {
		return err
	}

	if os.Getenv("DROP") == "y" {
		// insert static data
		db.Create(&PaymentMethods)
		db.Create(&Products)
	}

	return nil
}
