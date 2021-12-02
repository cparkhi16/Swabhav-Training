package main

import f "factory/autointerface"

type AutomobileFactory struct {
}
type Brand int

const (
	bmw      Brand = 1
	mercedes Brand = 2
	tesla    Brand = 3
)

func (a *AutomobileFactory) Make(b Brand) f.Automobile {
	if b == 1 {
		return NewBMW("BMW")
	} else if b == 2 {
		return NewMercedes("Mercedes")
	} else if b == 3 {
		return NewTesla("Tesla")
	}
	return nil
}
