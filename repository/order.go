package repository

import (
	"github.com/asrofilfachrulr/transaction-api/models/sql"
	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func (c *OrderRepo) Create(m *sql.Order) error {
	return c.DB.Create(m).Error
}
