package migrations

import "github.com/asrofilfachrulr/transaction-api/models/sql"

var PaymentMethods = []sql.PaymentMethod{
	{Name: 1, IsActive: true},
	{Name: 2, IsActive: false},
	{Name: 3, IsActive: true},
	{Name: 4, IsActive: false},
	{Name: 5, IsActive: true},
}

var Products = []sql.Product{
	{Name: 1, Price: 2000},
	{Name: 2, Price: 1000},
	{Name: 3, Price: 4000},
	{Name: 4, Price: 6000},
}
