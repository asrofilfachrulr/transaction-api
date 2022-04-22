package services

import (
	"github.com/asrofilfachrulr/transaction-api/models/sql"
	"github.com/asrofilfachrulr/transaction-api/repository"
)

type OrderService struct {
	OrderRepo        *repository.OrderRepo
	OrderPaymentRepo *repository.OrderPaymentRepo
	OrderItemRepo    *repository.OrderItemsRepo
}

func (or *OrderService) CreateOrder(o *sql.Order, oi *[]sql.OrderItem, opm *[]sql.OrderPayment) error {
	if err := or.OrderRepo.Create(o); err != nil {
		return err
	}

	for i := 0; i < len(*oi); i++ {
		(*oi)[i].OrderID = o.ID
	}

	for i := 0; i < len(*opm); i++ {
		(*opm)[i].OrderID = o.ID
	}

	if err := or.OrderItemRepo.Create(oi); err != nil {
		return err
	}

	if err := or.OrderPaymentRepo.Create(opm); err != nil {
		return err
	}

	return nil
}
