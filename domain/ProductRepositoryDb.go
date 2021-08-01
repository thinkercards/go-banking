package domain

import (
	"github.com/ashishjuyal/banking-lib/errs"
	"github.com/ashishjuyal/banking-lib/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ProductRepositoryDb struct {
	client *sqlx.DB
}

func (d ProductRepositoryDb) FindAll(status int) ([]Product, *errs.AppError) {
	var err error
	products := make([]Product, 0)

	if status == 0 {
		findAllSql :=
			`select		
		id
		,list_price
		,discontinued
		from products`

		err = d.client.Select(&products, findAllSql)
	} else {
		findAllSql :=
			`select 
			id
			,list_price
			,discontinued
			from products
		where discontinued = ?
		`
		err = d.client.Select(&products, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying products table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return products, nil
}

/*
func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers_banking where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}
*/

func NewProductRepositoryDb(dbClient *sqlx.DB) ProductRepositoryDb {
	return ProductRepositoryDb{dbClient}
}
