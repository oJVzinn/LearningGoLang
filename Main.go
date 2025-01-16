package main

import (
	"LearningGoLang/calculator"
	"LearningGoLang/countVowels"
	"LearningGoLang/operation"
	"LearningGoLang/primeNumber"
	"LearningGoLang/tasks"
	"log"
)

func main() {
	err := calculator.PrintTimeTable(10)
	catchError(err)

	err = calculator.Calculate(10, 20, operation.Sum)
	catchError(err)

	countVowels.CountVowels("Joao Victor")

	err = primeNumber.IsPrime(40)
	catchError(err)

	err = tasks.TaskSetup()
	catchError(err)
}

func catchError(error error) {
	if error != nil {
		log.Fatal(error)
	}
}
