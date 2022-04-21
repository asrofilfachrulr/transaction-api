package sql

type PaymentMethod struct {
	ID            uint           `gorm:"primaryKey;auto_increment"`
	Name          int            `gorm:"not null"`
	IsActive      bool           `gorm:"not null"`
	OrderPayments []OrderPayment `gorm:"constraint:OnDelete:CASCADE"`
}
