package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		isOk, err := customWithRand()

		if err != nil {
			fmt.Println("got error:", err)
		} else {
			fmt.Println("got ok:", isOk)
		}
	}
}

func customWithRand() (string, error) {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(101)
	isOk, err := custom(random)
	try err or string

	return isOk, nil
}

func custom(num int) (string, error) {
	if num > 50 {
		return "OK", nil
	}

	return "", errors.New("an error")
}
