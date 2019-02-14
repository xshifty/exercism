package triangle

import "math"

type Kind int

const (
	NaT Kind = iota + 1
	Equ
	Iso
	Sca
)

func IsNaT(a, b, c float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return true
	}

	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return true
	}

	if (a <= 0) || (b <= 0) || (c <= 0) {
		return true
	}

	if (a > b+c) || (b > a+c) || (c > b+a) {
		return true
	}

	return false
}

func IsEqu(a, b, c float64) bool {
	if a == b && a == c {
		return true
	}

	return false
}

func IsIso(a, b, c float64) bool {
	if a == b || a == c || b == c {
		return true
	}

	return false
}

func KindFromSides(a, b, c float64) Kind {
	if IsNaT(a, b, c) {
		return NaT
	}

	if IsEqu(a, b, c) {
		return Equ
	}

	if IsIso(a, b, c) {
		return Iso
	}

	return Sca
}
