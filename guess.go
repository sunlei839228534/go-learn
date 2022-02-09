package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	fmt.Println("我们生成了一个随机数1-100的随机数!")
	fmt.Println("你能猜一下它吗?")
	fmt.Println(target)

	success := false
	reader := bufio.NewReader(os.Stdin)
	for guess := 0; guess < 10; guess++ {
		fmt.Println("你还有", 10-guess, "次机会")
		fmt.Println("做一次猜测")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess > target {
			fmt.Println("猜高了!")
		} else if guess < target {
			fmt.Println("猜低了!")
		} else if guess == target {
			success = true
			fmt.Println("猜对啦")
			break
		}
	}

	if !success {
		fmt.Println("你失败了,最终答案是", target)
	}
}
