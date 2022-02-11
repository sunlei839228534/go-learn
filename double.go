package main

import (
	"fmt"
	"greeting"
)


func main () {
	greeting.Hello()
	amount := 6
	double(&amount)
	fmt.Println(amount)
}

func double(number *int) {
	*number *= 2
}	