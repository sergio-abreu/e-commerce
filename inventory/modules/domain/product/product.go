package product

type Product struct {
	Id           string `json:"id"`
	PriceInCents int    `json:"price_in_cents"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}
