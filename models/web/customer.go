package web

type (
	RegisterCustomerInput struct {
		Name    string `json:"name" binding:"required"`
		Address string `json:"address" binding:"required"`
	}
	RegisterCustomerOutput struct {
		CustomerID     uint   `json:"customer_id"`
		CustomerAddrID uint   `json:"customer_address_id"`
		Name           string `json:"name"`
		Address        string `json:"address"`
	}
	CustomerOutput struct {
		CustomerID     uint   `json:"customer_id"`
		CustomerAddrID uint   `json:"customer_address_id"`
		Name           string `json:"name" sql:"customer_name"`
	}
)
