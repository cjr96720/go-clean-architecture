package request

type AddProductRequest struct {
	ProductName string  `json:"product_name" binding:"required"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type UpdateProductRequest struct {
	ID          int     `json:"id" binding:"required"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}
