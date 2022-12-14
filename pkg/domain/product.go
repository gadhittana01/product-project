package domain

type Product struct {
	ID           int     `json:"id"`
	ProductName  string  `json:"product_name"`
	Qty          int     `json:"qty"`
	SellingPrice float64 `json:"selling_price"`
	PromoPrice   float64 `json:"promo_price"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}
