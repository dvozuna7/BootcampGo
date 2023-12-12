package domain

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int64   `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}
