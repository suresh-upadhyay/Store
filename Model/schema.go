package Model
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}


type Store struct {
	StoreId   int     `json:"store_id"`
	ProductId  string  `json:"product_id"`
	IsAvailable bool `json:"is_available"`
}

