package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m202408211711_autoname_machine_column() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202408211711",
		Migrate: func(db *gorm.DB) error {
			type Machine struct {
				AutoGenerateName bool `gorm:"default:true"`
			}

			return db.AutoMigrate(
				&Machine{},
			)
		},
		Rollback: nil,
	}
}
