package migrations

import modelSql "github.com/asrofilfachrulr/transaction-api/models/sql"

var Models = []any{
	&modelSql.Product{},
	&modelSql.Order{},
	&modelSql.OrderItem{},
	&modelSql.OrderPayment{},
	&modelSql.PaymentMethod{},
	&modelSql.Customer{},
	&modelSql.CustomerAddress{},
}
