package main

import (
	"log"
)

func main() {

	sum := 0
	current := 1
	next := 2

	for current < 4000000 {
		if current%2 == 0 {
			sum += current
		}
		fibonacci := current + next
		current = next
		next = fibonacci
	}

	log.Println(sum)
}
