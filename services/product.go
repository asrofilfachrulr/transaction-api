package services

import (
	"github.com/asrofilfachrulr/transaction-api/models/sql"
	"github.com/asrofilfachrulr/transaction-api/repository"
)

type ProductService struct {
	Repo *repository.ProductRepo
}

func (pms *ProductService) GetAllProduct(m *[]sql.Product) error {
	return pms.Repo.FindAll(m)
}
