package sql

type (
	Order struct {
		ID                uint            `gorm:"primaryKey;auto_increment"`
		CustomerAddressID uint            `gorm:"not null"`
		OrderItems        []OrderItem     `gorm:"constraint:OnDelete:CASCADE"`
		PaymentMethods    []PaymentMethod `gorm:"constraint:OnDelete:CASCADE"`
	}
	OrderItem struct {
		ID        uint `gorm:"primaryKey;auto_increment"`
		OrderID   uint `gorm:"not null"`
		ProductID uint `gorm:"not null"`
	}
	OrderPayment struct {
		ID        uint `gorm:"primaryKey;auto_increment"`
		OrderID   uint `gorm:"not null"`
		PaymentID uint `gorm:"not null"`
	}
)
