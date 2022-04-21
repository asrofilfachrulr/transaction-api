package migrations

import modelSql "github.com/asrofilfachrulr/transaction-api/models/sql"

var Models = []any{
	&modelSql.Customer{},
	&modelSql.CustomerAddress{},
	&modelSql.PaymentMethod{},
	&modelSql.Product{},
}
