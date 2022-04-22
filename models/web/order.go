package web

type (
	PostOrderInput struct {
		CustomerAddrID   uint   `json:"customer_address" binding:"required"`
		PaymentMethodIDs []uint `json:"payment_methods" binding:"required"`
		ProductIDs       []uint `json:"products" binding:"required"`
	}
	PostOrderOutput struct {
		OrderID uint `json:"order_id"`
	}
)
