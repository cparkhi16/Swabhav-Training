package main

import (
	"fmt"
	"sync"
)

// Observer interface
type Observer interface {
	Update(price float64)
}

// Subject struct
type Subject struct {
	observers map[Observer]struct{}
	mu        sync.Mutex
}

func (s *Subject) Attach(o Observer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.observers[o] = struct{}{}
}

func (s *Subject) Detach(o Observer) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.observers, o)
}

func (s *Subject) Notify(price float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for o := range s.observers {
		o.Update(price)
	}
}

// Stock struct
type Stock struct {
	Subject
	price float64
}

func NewStock() *Stock {
	return &Stock{
		Subject: Subject{observers: make(map[Observer]struct{})},
	}
}

func (s *Stock) SetPrice(price float64) {
	s.price = price
	s.Notify(price)
}

// Investor struct
type Investor struct {
	name  string
	stock *Stock
}

func (i *Investor) Update(price float64) {
	fmt.Printf("Investor %s notified of stock price: %.2f\n", i.name, i.stock.price)
}

func main() {
	stock := NewStock()

	investor1 := &Investor{name: "Alice", stock: stock}
	investor2 := &Investor{name: "Bob", stock: stock}

	stock.Attach(investor1)
	stock.Attach(investor2)

	stock.SetPrice(101.12)
	stock.SetPrice(104.31)

	stock.Detach(investor1)
	stock.SetPrice(150.9)
}
