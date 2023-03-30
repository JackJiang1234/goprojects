package factory

import "fmt"

type Country interface {
	Greet()
}

type country struct {
	name string
}

func (c country) Greet(){
	fmt.Printf("hi %s", c.name)
}

func NewCountry(cname string) Country{
	return &country{
		name: cname,
	}
}