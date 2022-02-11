package main

import "fmt"

func main() {
	var sum float64 = 0
	 numbers := [3]float64{71.8,56.2,89.5}
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)
}