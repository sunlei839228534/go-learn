package main

import (
	"time"
	"fmt"
)

func main() {
	var dates [3]time.Time
	dates[0] = time.Unix(1257894000 ,0)
	dates[1] = time.Unix(1447894000 ,0)
	dates[2] = time.Unix(1507894000 ,0)
	for index , value := range dates {
		fmt.Println(index, value)
	}
}