package dto

type OrderRequest struct {
	ProductName string `json:"product_name" valo:"notblank"`
	CustomerID  int32  `json:"-"`
	Username    string `json:"-"`
}
