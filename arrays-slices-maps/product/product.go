package product

type Product struct {
	title string
	id    int
	price float64
}

func New(title string, id int, price float64) Product {
	return Product{title, id, price}
}
