package raindrops

import (
	"fmt"
)

func Convert(number int) string {
	if number < 3 {
		return fmt.Sprintf("%d", number)
	}

	result := ""
	factors := []int{3, 5, 7}
	replacements := []string{"Pling", "Plang", "Plong"}

	for n, f := range factors {
		if number < f {
			break
		}

		if number%f == 0 {
			result = fmt.Sprintf("%s%s", result, replacements[n])
		}
	}

	if len(result) == 0 {
		result = fmt.Sprintf("%d", number)
	}

	return result
}
