package algo

import "fmt"

type Fifo struct {
}

func NewFIFO() *Fifo {
	return &Fifo{}
}
func (l *Fifo) Evict(c *Cache) {
	//fmt.Println(c.fifoQueue)
	c.Remove(c.Queue[0])
	//fmt.Println(c.storage)
	fmt.Println("Evicting by fifo strtegy")
}
