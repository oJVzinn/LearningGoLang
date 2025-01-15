package main

import (
	"LearningGoLang/calculator"
	"LearningGoLang/operation"
	"log"
)

func main() {
	calculator.PrintTimeTable(5)
	err := calculator.Calculate(10, 20, operation.Sum)
	if err != nil {
		log.Fatal(err)
	}

}
