package stringSorted

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func Sorted(words []string) error {
	if words == nil || len(words) == 0 {
		return errors.New("no words provided")
	}

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	builder := strings.Builder{}
	for i := 0; i < len(words); i++ {
		builder.WriteString(words[i])
		if i != len(words)-1 {
			builder.WriteString(", ")
		}
	}

	fmt.Println("The final sorted: " + builder.String())
	return nil
}
