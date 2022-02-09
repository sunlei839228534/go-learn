package main

import (
	"fmt"
	"log"
)

func main() {
	var amount, total float64
	amount, err := paintNeeded(4.2, 3.0)
	if err != nil {
		log.Fatal(err)
	}
	total += amount
	amount, err = paintNeeded(12, 3.5)
	if err != nil {
		log.Fatal(err)
	}
	total += amount
	fmt.Printf("总共需要: %0.2f 升\n", total)
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("宽度 %0.2f 无效", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("高度 %0.2f 无效", height)
	}
	area := width * height
	return area / 10.08, nil
}
