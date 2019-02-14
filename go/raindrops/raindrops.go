package raindrops

import (
	"fmt"
)

func Convert(number int) string {
	if number < 3 {
		return fmt.Sprintf("%d", number)
	}

	result := ""
	replacements := []struct {
		factor int
		word   string
	}{
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"}}

	for _, r := range replacements {
		if number < r.factor {
			break
		}

		if (number % r.factor) == 0 {
			result = fmt.Sprintf("%s%s", result, r.word)
		}
	}

	if len(result) == 0 {
		result = fmt.Sprintf("%d", number)
	}

	return result
}
