package dto

type OrderRequest struct {
	ProductName string `json:"product_name" valo:"notblank"`
	CustomerID  string `json:"-"`
	Username    string `json:"-"`
}
