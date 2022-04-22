package services

import (
	"github.com/asrofilfachrulr/transaction-api/models/sql"
	"github.com/asrofilfachrulr/transaction-api/repository"
)

type PaymentMethodService struct {
	Repo *repository.PaymentMethodRepo
}

func (pms *PaymentMethodService) GetAllMethod(m *[]sql.PaymentMethod) error {
	return pms.Repo.FindAll(m)
}
