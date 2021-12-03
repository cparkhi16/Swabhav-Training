package algo

import "fmt"

type Lfu struct {
}

func (l *Lfu) Evict(c *Cache) {
	fmt.Println("Evicting by lfu strtegy")
	min := 500
	lfuElement := ""
	for k, v := range c.freq {
		if v < min {
			lfuElement = k
			min = v
		}
	}
	if lfuElement == "" {
		lfuElement = c.Queue[0]
	}
	c.Remove(lfuElement)
	index := 0
	for i, v := range c.Queue {
		if v == lfuElement {
			index = i
			break
		}
	}
	c.Queue = append(c.Queue[:index], c.Queue[index+1:]...)

}
