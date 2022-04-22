package repository

import (
	"github.com/asrofilfachrulr/transaction-api/models/sql"
	"gorm.io/gorm"
)

type PaymentMethodRepo struct {
	DB *gorm.DB
}

func (c *PaymentMethodRepo) FindAll(m *[]sql.PaymentMethod) error {
	return c.DB.Find(m).Error
}
