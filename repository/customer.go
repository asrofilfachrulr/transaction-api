package repository

import (
	"github.com/asrofilfachrulr/transaction-api/models/sql"
	"gorm.io/gorm"
)

type CustRepo struct {
	DB *gorm.DB
}

func (c *CustRepo) Create(m *sql.Customer) error {
	return c.DB.Create(m).Error
}
