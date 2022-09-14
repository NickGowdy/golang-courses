package product

type Product struct {
	ProductID      int    `json:"productId"`
	Manafacturer   string `json:"manafacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProduceName    string `json:"productName"`
}
