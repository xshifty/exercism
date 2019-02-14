package leap

func IsLeapYear(year int) bool {
	isLeap := (year % 4) == 0
	isLeap = isLeap && ((year % 100) != 0)
	isLeap = isLeap || ((year % 400) == 0)

	return isLeap
}
