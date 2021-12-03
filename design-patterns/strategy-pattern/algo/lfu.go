package algo

import "fmt"

type Lfu struct {
}

func (l *Lfu) Evict(c *Cache) {
	fmt.Println("Evicting by lfu strategy")
	min := 500
	lfuElement := ""
	for k, v := range c.freq {
		if v < min {
			lfuElement = k
			min = v
		}
	}
	if lfuElement == "" {
		lfuElement = c.queue[0]
	}
	c.Remove(lfuElement)
	index := 0
	for i, v := range c.queue {
		if v == lfuElement {
			index = i
			break
		}
	}
	delete(c.freq, lfuElement)
	c.queue = append(c.queue[:index], c.queue[index+1:]...)

}
