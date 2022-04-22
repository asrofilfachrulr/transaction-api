package repository

import (
	"github.com/asrofilfachrulr/transaction-api/models/sql"
	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

func (c *ProductRepo) FindAll(m *[]sql.Product) error {
	return c.DB.Find(m).Error
}
