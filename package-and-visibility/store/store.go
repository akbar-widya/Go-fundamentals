package store

type Product struct {
	Name string
	price float64
}

func NewProduct(name string, price float64) Product {
	return Product{Name: name, price: price}
}