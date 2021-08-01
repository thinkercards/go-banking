package dto

type ProductResponse struct {
	SupplierId             string  `json:"supplier_ids"`
	Id                     int     `json:"id"`
	ProductCode            string  `json:"product_code"`
	ProductName            string  `json:"product_name"`
	Description            string  `json:"description"`
	StandardCost           float32 `json:"standard_cost"`
	ListPrice              float32 `json:"list_price"`
	ReorderLevel           int     `json:"reorder_level"`
	TargetLevel            int     `json:"target_level"`
	QuantityPerUnit        string  `json:"quantity_per_unit"`
	Discontinued           string  `json:"discontinued"`
	MinimumReorderQuantity int     `json:"minimum_reorder_quantity"`
	Category               string  `json:"category"`
}
