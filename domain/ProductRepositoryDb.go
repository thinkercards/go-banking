package domain

import (
	"github.com/ashishjuyal/banking-lib/errs"
	"github.com/ashishjuyal/banking-lib/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
/*

https://stackoverflow.com/questions/44891030/scan-error-unsupported-scan-storing-driver-value-type-nil-into-type-string
You can use any of the below two solutions:-

You can use sql.NullString to handle the field before using scan(). OR
You can replace all the possible NULL values with the desired string say '' from the query itself.
For implementing the 1st solution refer to the @RayfenWindspear answer. For the 2nd solution update the query as below:-

SELECT colm1, colm2, COALESCE(photo, '') photo, colm4 FROM Article WHERE photo IS NULL
For MySQL use IFNULL() or COALESCE() function to return an alternative value if an expression is NULL:

https://www.w3schools.com/sql/sql_isnull.asp

*/
type ProductRepositoryDb struct {
	client *sqlx.DB
}

func (d ProductRepositoryDb) FindAll(status int) ([]Product, *errs.AppError) {
	var err error
	products := make([]Product, 0)

	if status == 0 {
		findAllSql :=
			`select

			COALESCE(supplier_ids, '') as supplier_ids
			,id
			,COALESCE(product_code, '') as product_code
			,COALESCE(product_name, '') as product_name
			,standard_cost
			,list_price
			,COALESCE(reorder_level, 0) as reorder_level
			,COALESCE(target_level, 0) as target_level
			,COALESCE(quantity_per_unit, '') as quantity_per_unit
			,COALESCE(description, '') as description
			,discontinued
			,COALESCE(minimum_reorder_quantity, 0) as minimum_reorder_quantity
			,COALESCE(category, '') as category
		    from products`

		err = d.client.Select(&products, findAllSql)
	} else {
		findAllSql :=
			`select 
			COALESCE(supplier_ids, '') as supplier_ids
			,id
			,COALESCE(product_code, '') as product_code
			,COALESCE(product_name, '') as product_name
			,standard_cost
			,list_price
			,COALESCE(reorder_level, 0) as reorder_level
			,COALESCE(target_level, 0) as target_level
			,COALESCE(quantity_per_unit, '') as quantity_per_unit
			,COALESCE(description, '') as description
			,discontinued
			,COALESCE(minimum_reorder_quantity, 0) as minimum_reorder_quantity
			,COALESCE(category, '') as category
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
