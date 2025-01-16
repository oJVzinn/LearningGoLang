package primeNumber

import (
	"errors"
	"fmt"
)

func IsPrime(number int) error {
	if number < 2 {
		return errors.New("the number is invalid")
	}

	isPrime := true
	if number > 2 {
		if number%2 == 0 || number%3 == 0 || number%5 == 0 || number%7 == 0 {
			isPrime = false
		}
	}

	message := func() string {
		if isPrime {
			return fmt.Sprintf("The number %d, is prime", number)
		}

		return fmt.Sprintf("The number %d, is not a prime", number)
	}

	fmt.Println(message())
	return nil
}
