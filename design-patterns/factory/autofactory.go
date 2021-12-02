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

func (b Brand) String() string {
	return [...]string{"bmw", "mercedes", "tesla"}[b-1]
}
func (a *AutomobileFactory) Make(b Brand) f.Automobile {
	if b.String() == "bmw" {
		return &BMW{modelName: "BMW"}
	} else if b.String() == "mercedes" {
		return &Mercedes{modelName: "Mercedes"}
	} else if b.String() == "tesla" {
		return &Tesla{modelName: "Tesla"}
	}
	return nil
}
