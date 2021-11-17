package main

import "fmt"

func main() {
	jewelry := make(map[string]float64)

	jewelry["necklace"] = 89.99
	jewelry["earrings"] = 79.99

	clothing := map[string]float64{"pants": 59.99, "shirt": 39.99}

	fmt.Println("Earrings:", jewelry["earrings"])
	fmt.Println("Necklaces", jewelry["necklace"])
	fmt.Println("Shirt:", clothing["shirt"])
	fmt.Println("Pants", clothing["pants"])

	data := []string{"a", "c", "e", "a", "e"}
	counts := make(map[string]int)
	for _, item := range data {
		counts[item]++
	}

	letters := []string{"a", "b", "c", "d", "e"}

	for _, letter := range letters {
		count, ok := counts[letter]
		if !ok {
			fmt.Printf("%s: not found\n", letter)
		} else {
			fmt.Printf("%s: %d\n", letter, count)

		}
	}
}
