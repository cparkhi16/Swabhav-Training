package algo

import "fmt"

type Fifo struct {
}

func NewFIFO() *Fifo {
	return &Fifo{}
}

func (l *Fifo) Evict(c *Cache) {
	//fmt.Println(c.fifoQueue)
	c.Remove(c.queue[0])
	c.queue = c.queue[1:]
	//fmt.Println("Queue", c.Queue)
	//fmt.Println(c.storage)
	fmt.Println("Evicting by fifo strategy")
}
