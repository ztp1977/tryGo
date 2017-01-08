package http

import "errors"

type Cart struct {
	products []string
}

func New() *Cart {
	c := new(Cart)
	c.products = make([]string, 0)
	return c
}

func (c *Cart) Add(s string) {
	c.products = append(c.products, s)
}

func (c *Cart) GetAll() []string {
	return c.products
}

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数は0以外でなければなりません")
	}
	return a / b, nil
}
