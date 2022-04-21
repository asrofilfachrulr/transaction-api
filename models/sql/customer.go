package sql

type (
	Customer struct {
		ID                uint              `gorm:"primaryKey;auto_increment"`
		CustomerName      string            `gorm:"type:varchar(50);not null"`
		CustomerAddresses []CustomerAddress `gorm:"constraint:OnDelete:CASCADE"`
	}
	CustomerAddress struct {
		ID         uint    `gorm:"primaryKey;auto_increment"`
		Address    string  `gorm:"not null"`
		CustomerID uint    `gorm:"not null"`
		Orders     []Order `gorm:"constraint:OnDelete:CASCADE"`
	}
)
