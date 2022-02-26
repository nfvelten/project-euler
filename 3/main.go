package main

import (
	"log"
)

func main() {
	n := 600851475143
	i := 2

	for i < n {
		for n%i == 0 {
			log.Println("The value of i is ", i)
			n = n / i
			log.Println("The value of n is ", n)
		}
		i++
	}
	log.Println(i)
}
