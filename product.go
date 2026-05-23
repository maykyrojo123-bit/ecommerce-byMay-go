package product

import "errors"

type Product struct {
	id    string
	Name  string
	Price float64
	stock int
}

func NewProduct(id string, name string, price float64, stock int) Product {
	return Product{id: id, Name: name, Price: price, stock: stock}
}

func (p *Product) GetID() string {
	return p.id
}

func (p *Product) GetStock() int {
	return p.stock
}

func (p *Product) ReduceStock(quantity int) error {
	if quantity > p.stock {
		return errors.New("insufficient stock available")
	}
	p.stock -= quantity
	return nil
}
