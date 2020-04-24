package main

/*
	Generate random answer 'Yes' or 'No'
*/

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(3)
	num := rand.Intn(3)
	if num == 1 {
		fmt.Println("Yes")
	}
	if num == 2 {
		fmt.Println("No")
	}

}
