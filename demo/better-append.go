package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	@numbers <<< 4
	fmt.Println(numbers)
}
