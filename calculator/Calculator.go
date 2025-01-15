package calculator

import (
	"LearningGoLang/operation"
	"errors"
	"fmt"
)

func PrintTimeTable(number int) {
	fmt.Println("TimeTable x", number)
	for i := 1; i <= 10; i++ {
		finalValue := number * i
		fmt.Println("value for", i, "*", number, "=", finalValue)
	}
}

func Calculate(numberOne int, numberTwo int, op operation.FactorOfOperation) error {
	switch op {
	case operation.Sum:
		result := numberOne + numberTwo
		fmt.Println("Result:", result)
		return nil

	case operation.Sub:
		result := numberOne - numberTwo
		fmt.Println("Result:", result)
		return nil

	case operation.Division:
		result := numberOne / numberTwo
		fmt.Println("Result:", result)
		return nil

	case operation.Multiply:
		result := numberOne * numberTwo
		fmt.Println("Result:", result)
		return nil

	default:
		return errors.New("Invalid operation")
	}
}
