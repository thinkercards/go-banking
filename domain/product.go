package domain

import (
	"github.com/ashishjuyal/banking-lib/errs"
	"github.com/ashishjuyal/banking/dto"
)

type Product struct {
	//SupplierId  string `db:"supplier_ids"`
	Id int `db:"id"`
	//ProductCode string `db:"product_code"`
	//ProductName string `db:"product_name"`
	//Description            sql.NullString `db:"description"`
	//StandardCost           float32 `db:"standard_cost"`
	ListPrice float32 `db:"list_price"`
	//ReorderLevel           int     `db:"reorder_level"`
	//TargetLevel            int     `db:"target_level"`
	//QuantityPerUnit        string  `db:"quantity_per_unit"`
	Discontinued int `db:"discontinued"`
	//MinimumReorderQuantity int     `db:"minimum_reorder_quantity"`
	//Category               string  `db:"category"`
	//Attachments longblob  `db:"attachments"`

}

func (p Product) discontinuedAsText() string {
	statusAsText := "yes"
	if p.Discontinued == 0 {
		statusAsText = "no"
	}
	return statusAsText
}

func (p Product) ToDto() dto.ProductResponse {
	return dto.ProductResponse{
		//	SupplierId:  p.SupplierId,
		Id: p.Id,
		//	ProductCode: p.ProductCode,
		//	ProductName: p.ProductName,
		//Description: p.Description,
		//	StandardCost:           p.StandardCost,
		ListPrice: p.ListPrice,
		//	ReorderLevel:           p.ReorderLevel,
		//	TargetLevel:            p.TargetLevel,
		//	QuantityPerUnit:        p.QuantityPerUnit,
		Discontinued: p.discontinuedAsText(),
		//	MinimumReorderQuantity: p.MinimumReorderQuantity,
		//	Category:               p.Category,
	}
}

type ProductRepository interface {
	FindAll(status int) ([]Product, *errs.AppError)
	/*ById(string) (*Customer, *errs.AppError)*/
}
