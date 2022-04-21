package sql

type Order struct {
	ID                uint           `gorm:"primaryKey;auto_increment"`
	CustomerAddressID uint           `gorm:"not null"`
	OrderItems        []OrderItem    `gorm:"constraint:OnDelete:CASCADE"`
	OrderPayments     []OrderPayment `gorm:"constraint:OnDelete:CASCADE"`
}
type OrderItem struct {
	ID        uint `gorm:"primaryKey;auto_increment"`
	OrderID   uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
}
type OrderPayment struct {
	ID              uint `gorm:"primaryKey;auto_increment"`
	OrderID         uint `gorm:"not null"`
	PaymentMethodID uint `gorm:"not null"`
}
