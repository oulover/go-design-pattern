package simplefactory

import (
	"fmt"
	"sync"
)

type singleTon struct {
}

func (s *singleTon) Show() {
	fmt.Println("Hello Im singleton")
}

var (
	once   sync.Once
	single *singleTon
)

func GetSingleInstance() *singleTon {
	once.Do(func() {
		single = &singleTon{}
	})
	return single
}
