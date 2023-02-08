package simplefactory

import "fmt"

// Car is an interface
type Car interface {
	// Run : car running with wheels
	Run(wheel int) string
}

// coachCar is one kind of car
type coachCar struct {
}

func (c *coachCar) Run(wheel int) string {
	return fmt.Sprintf("I am Coach,Running with %d wheels", wheel)
}

type trainCar struct {
}

func (t *trainCar) Run(wheel int) string {
	return fmt.Sprintf("I am train,Running with %d wheels", wheel)
}

func NewCar(i int) Car {
	if i == 1 {
		return &coachCar{}
	} else if i == 2 {
		return &trainCar{}
	} else {
		return nil
	}
}
