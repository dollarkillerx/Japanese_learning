package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := []string{
		"あ", "い", "う", "え", "お",
		"ァ", "イ", "ウ", "エ", "オ",
	}

	var i string
	for {
		rand.Seed(time.Now().UnixNano())
		fmt.Println(r[rand.Intn(len(r))])

		fmt.Scanln(&i)
	}
}
