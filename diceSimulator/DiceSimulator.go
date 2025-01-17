package diceSimulator

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
)

func LaunchDiceSimulator(totalRounds int) error {
	if totalRounds < 1 || totalRounds > 10 {
		return errors.New("total round must be between 1 and 10")
	}

	builder := strings.Builder{}
	for i := 0; i < totalRounds; i++ {
		result := rand.IntN(6)
		builder.WriteString("Result for Round " + strconv.Itoa(i+1) + ": " + strconv.Itoa(result) + "\n")
	}

	fmt.Println(builder.String())
	return nil
}
