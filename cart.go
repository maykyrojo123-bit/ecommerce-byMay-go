package cart

import "ecommerce-poo/product"

type CartItem struct {
	Product  product.Product
	Quantity int
}

type Cart struct {
	Items []CartItem
}

func (c *Cart) AddItem(p product.Product, qty int) {
	item := CartItem{Product: p, Quantity: qty}
	c.Items = append(c.Items, item)
}

func (c *Cart) CalculateSubtotal() float64 {
	var subtotal float64
	for _, item := range c.Items {
		subtotal += item.Product.Price * float64(item.Quantity)
	}
	return subtotal
}
