package bulletin

import (
	"errors"
	"fmt"
)

func CalculateAverage(grades []int) error {
	if len(grades) != 4 {
		return errors.New("the grades is incorrect")
	}

	var sumGrades int
	for _, grade := range grades {
		sumGrades += grade
	}

	fmt.Println("The final average grade is", sumGrades/len(grades))
	return nil
}
