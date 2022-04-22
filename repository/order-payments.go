package repository

import (
	"github.com/asrofilfachrulr/transaction-api/models/sql"
	"gorm.io/gorm"
)

type OrderPaymentRepo struct {
	DB *gorm.DB
}

func (c *OrderPaymentRepo) Create(m *[]sql.OrderPayment) error {
	return c.DB.Create(m).Error
}
