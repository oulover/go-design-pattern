package simplefactory

import (
	"fmt"
	"testing"
)

func TestCreateCar1(t *testing.T) {
	car := NewCar(1)
	run := car.Run(4)
	fmt.Println(run)
}

func TestCreateCar2(t *testing.T) {
	car := NewCar(2)
	run := car.Run(100)
	fmt.Println(run)
}
