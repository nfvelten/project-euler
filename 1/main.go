package main

import "log"

func main() {
	number := 1000
	sum := 0

	for i := 0; i < number; i++ {
		result := isMultipleOfThreeOrFive(i)
		if result {
			sum += i
		}
	}

	log.Println(sum)
}

func isMultipleOfThreeOrFive(number int) bool {
	if number%3 == 0 || number%5 == 0 {
		return true
	}
	return false
}
