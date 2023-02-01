package foobar

type Car struct {
	Miles int
	Year  int
	Model string
	Maker string
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

func NewCar(miles, year int, model, maker string) *Car {
	return &Car{
		Miles: miles,
		Year:  year,
		Model: model,
		Maker: maker,
	}
}
