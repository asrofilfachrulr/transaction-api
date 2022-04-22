package sql

type Product struct {
	ID         uint        `json:"id" gorm:"primaryKey;auto_increment"`
	Name       int         `json:"name" gorm:"not null"`
	Price      float64     `json:"price" gorm:"not null"`
	OrderItems []OrderItem `json:",omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:ProductID"`
}
