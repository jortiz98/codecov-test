package foobar_test

import (
	"foobar"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testcase = []struct {
	name string
	car  *foobar.Car
}{
	{name: "BMW", car: &foobar.Car{
		Miles: 12951,
		Year:  1985,
		Model: "M3",
		Maker: "BMW",
	}},
	{name: "Audi", car: &foobar.Car{
		Miles: 291000,
		Year:  2001,
		Model: "r8",
		Maker: "Audi",
	}},
	{name: "Tesla", car: &foobar.Car{
		Miles: 6540,
		Year:  2022,
		Model: "Model Y",
		Maker: "Tesla",
	}},
}

func TestCar_GetMaker(t *testing.T) {
	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.car.GetMaker()

			assert.Equal(t, tc.car.Maker, actual)
		})
	}

}

func TestCar_GetModel(t *testing.T) {
	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.car.GetModel()

			assert.Equal(t, tc.car.Model, actual)
		})
	}

}

func TestCar_GetMiles(t *testing.T) {
	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.car.GetMiles()

			assert.Equal(t, tc.car.Miles, actual)
		})
	}
}

//func TestCar_GetYear(t *testing.T) {
//	for _, tc := range testcase {
//		t.Run(tc.name, func(t *testing.T) {
//			actual := tc.car.GetYear()
//
//			assert.Equal(t, tc.car.Year, actual)
//		})
//	}
//}
