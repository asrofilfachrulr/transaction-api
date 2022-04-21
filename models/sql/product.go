package sql

type Product struct {
	ID         uint        `gorm:"primaryKey;auto_increment"`
	Name       int         `gorm:"not null"`
	Price      float64     `gorm:"not null"`
	OrderItems []OrderItem `gorm:"constraint:OnDelete:CASCADE"`
}
