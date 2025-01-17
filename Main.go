package main

import (
	"LearningGoLang/bulletin"
	"LearningGoLang/calculator"
	"LearningGoLang/countVowels"
	"LearningGoLang/diceSimulator"
	"LearningGoLang/operation"
	"LearningGoLang/primeNumber"
	"LearningGoLang/stringSorted"
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

	err = bulletin.CalculateAverage([]int{8, 3, 10, 4})
	catchError(err)

	err = stringSorted.Sorted([]string{"pao", "chuveiro", "oi", "c", "Bleybleide", "joao victor"})
	catchError(err)

	err = diceSimulator.LaunchDiceSimulator(8)
	catchError(err)
}

func catchError(error error) {
	if error != nil {
		log.Fatal(error)
	}
}
