package repository

import (
	"github.com/asrofilfachrulr/transaction-api/models/sql"
	"gorm.io/gorm"
)

type CustAddrRepo struct {
	DB *gorm.DB
}

func (c *CustAddrRepo) Create(m *sql.CustomerAddress) error {
	return c.DB.Create(m).Error
}
