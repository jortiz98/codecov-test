package foobar_test

import (
	"foobar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCar_GetMaker(t *testing.T) {
	testcase := []struct {
		name string
		car  *foobar.Car
	}{
		{name: "BMW", car: &foobar.Car{
			Miles: 0,
			Year:  0,
			Model: "M3",
			Maker: "BMW",
		}},
		{name: "Audi", car: &foobar.Car{
			Miles: 0,
			Year:  0,
			Model: "r8",
			Maker: "Audi",
		}},
		{name: "Tesla", car: &foobar.Car{
			Miles: 0,
			Year:  0,
			Model: "Model Y",
			Maker: "Tesla",
		}},
	}

	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.car.GetMaker()

			assert.Equal(t, tc.car.Maker, actual)
		})
	}

}

func TestCar_GetModel(t *testing.T) {
	testcase := []struct {
		name string
		car  *foobar.Car
	}{
		{name: "BMW", car: &foobar.Car{
			Miles: 0,
			Year:  0,
			Model: "M3",
			Maker: "BMW",
		}},
		{name: "Audi", car: &foobar.Car{
			Miles: 0,
			Year:  0,
			Model: "r8",
			Maker: "Audi",
		}},
		{name: "Tesla", car: &foobar.Car{
			Miles: 0,
			Year:  0,
			Model: "Model Y",
			Maker: "Tesla",
		}},
	}

	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.car.GetModel()

			assert.Equal(t, tc.car.Model, actual)
		})
	}

}
