package repository

import (
	"github.com/asrofilfachrulr/transaction-api/models/sql"
	"gorm.io/gorm"
)

type OrderItemsRepo struct {
	DB *gorm.DB
}

func (c *OrderItemsRepo) Create(m *[]sql.OrderItem) error {
	return c.DB.Create(m).Error
}
