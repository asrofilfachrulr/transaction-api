package sql

type PaymentMethod struct {
	ID            uint           `json:"id" gorm:"primaryKey;auto_increment"`
	Name          int            `json:"name" gorm:"not null"`
	IsActive      bool           `json:"is_active" gorm:"not null"`
	OrderPayments []OrderPayment `json:",omitempty" gorm:"constraint:OnDelete:CASCADE"`
}
