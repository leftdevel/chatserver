package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(keys(map[string]struct{}{
	// 	"dog": struct{}{},
	// 	"cat": struct{}{},
	// }))

	fmt.Println("Hello world")

	var a = 1
	fmt.Println(a)

	message := "hola"
	fmt.Println(message)

	var b, c int = 2, 3

	fmt.Println(b)
	fmt.Println(c)

	var d [5]int

	fmt.Println(d)
	// d = append(d, 1)
	fmt.Println(d)

	numbers := make([]int, 5, 10)
	fmt.Println(numbers)

	numbers2 := []int{2, 4, 6, 8}
	var max int = int(math.Min(float64(len(numbers2)), 10))

	fmt.Println(numbers2[2:max])
}

func keys(m map[string]struct{}) []string {
	ret := make([]string, len(m))
	i := 0
	for key := range m {
		ret[i] = key
		i++
	}
	return ret
}
