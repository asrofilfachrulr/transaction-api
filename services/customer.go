package services

import (
	"github.com/asrofilfachrulr/transaction-api/models/sql"
	"github.com/asrofilfachrulr/transaction-api/repository"
)

type CustomerService struct {
	CustRepo     *repository.CustRepo
	CustAddrRepo *repository.CustAddrRepo
}

func (cs *CustomerService) CreateCustomer(c *sql.Customer, ca *sql.CustomerAddress) error {
	if err := cs.CustRepo.Create(c); err != nil {
		return err
	}

	ca.CustomerID = c.ID

	if err := cs.CustAddrRepo.Create(ca); err != nil {
		return err
	}

	return nil
}

func (cs *CustomerService) GetAllCustomer(c *[]sql.Customer) error {
	if err := cs.CustRepo.FindAll(c); err != nil {
		return err
	}
	return nil
}
