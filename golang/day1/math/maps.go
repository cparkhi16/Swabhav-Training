package math

import "fmt"

func MapTest() {
	var a map[string]int //a nil map can't assign values afterwards
	fmt.Println(" map not initialized ", a)
	//a["st"] = 1 //assignment to entry in nil map
	//fmt.Println(" map after initialized ", a)
	b := map[string]int{
		"chinmay": 1,
		"RAJ":     2,
	}
	fmt.Println(" normal map init ", b)

	c := make(map[string]int)
	c["test"] = 1
	fmt.Println(" NEW Map using make ", c)
	c = b
	b["RAJ"] = 100
	fmt.Println("after manipulation b and c ", b, c)
	// maps are copied by reference
	delete(b, "RAJ")
	fmt.Println("after deletion b and c ", b, c)
	c["name"] = 4
	c["llm"] = 1
	val, ok := c["llm"]
	if ok {
		fmt.Println(" val accessed for key llm is ", val)
	} else {
		fmt.Println(" no key llm found ")
	}
	for k, v := range c {
		fmt.Println(" key in c ", k, " val in c ", v)
	}
}

func ComplexMap() {
	c := make(map[string]map[string]int)
	c["k"] = map[string]int{
		"v": 1,
	}
	fmt.Println(" complex map is ", c)

	v := make(map[string]map[string]int)

	for k, vl := range c {
		for kl, vll := range vl {
			v[k] = map[string]int{kl: vll}
		}
	}
	fmt.Println(" deep copy a map ", v)
	c["k"] = map[string]int{
		"va": 1,
	}
	fmt.Println("after deep copy a map v is  ", v, " og map is ", c)
}
