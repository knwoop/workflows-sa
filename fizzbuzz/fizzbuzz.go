package fizzbuzz

import (
	"strconv"
)

func Fizzbuzz(val int) string {
	if val%15 == 0 {
		return "FizzBuzz"
	}

	if val%3 == 0 {
		return "Fizz"
	}

	if val%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(val)
}
