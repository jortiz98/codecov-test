package foobar

type Car struct {
	Miles int
	Year  int
	Model string
	Maker string
	Color string
	Price float64
}

func (c *Car) GetYear() int {
	return c.Year
}

func (c *Car) GetMiles() int {
	return c.Miles
}

func (c *Car) GetModel() string {
	return c.Model
}

func (c *Car) GetMaker() string {
	return c.Maker
}

func (c *Car) GetColor() string {
	return c.Color
}

func (c *Car) GetPrice() float64 {
	return c.Price
}
