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
		return &BMW{modelName: "BMW"}
	} else if b == 2 {
		return &Mercedes{modelName: "Mercedes"}
	} else if b == 3 {
		return &Tesla{modelName: "Tesla"}
	}
	return nil
}
