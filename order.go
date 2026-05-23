package order

import (
	"ecommerce-poo/cart"
	"time"
)

type TaxCalculator interface {
	CalculateTax(amount float64) float64
}

type EcuadorTax struct {
	Rate float64
}

func (et EcuadorTax) CalculateTax(amount float64) float64 {
	return amount * et.Rate
}

type Order struct {
	ID        string
	CartItems []cart.CartItem
	Subtotal  float64
	Tax       float64
	Total     float64
	CreatedAt time.Time
}

func (o *Order) ProcessCheckout(id string, c cart.Cart, taxCalc TaxCalculator) {
	o.ID = id
	o.CartItems = c.Items
	o.Subtotal = c.CalculateSubtotal()
	o.Tax = taxCalc.CalculateTax(o.Subtotal)
	o.Total = o.Subtotal + o.Tax
	o.CreatedAt = time.Now()
}
