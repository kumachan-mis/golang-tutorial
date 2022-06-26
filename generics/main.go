package main

import "fmt"

type Number interface {
	int32 | float64
}

func sumOfNumverValues[K comparable, V Number](numberMap map[K]V) V {
	var sum V = V(0)
	for _, value := range numberMap {
		sum += value
	}
	return sum
}

func main() {
	fmt.Println(sumOfNumverValues(map[string]int32{"fst": 10, "snd": 20}))
	fmt.Println(sumOfNumverValues(map[string]float64{"fst": 1.2, "snd": 6.7}))
}
