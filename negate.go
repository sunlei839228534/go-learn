package main

import "fmt"

func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
}

func negate(myBoolean *bool)  {
	*myBoolean = !*myBoolean
}