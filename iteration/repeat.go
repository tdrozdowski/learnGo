package iteration

import (
	"math"
	"strconv"
)

func Repeat(what string, howMany int) string {
	var repeated string
	for i := 0; i < howMany; i++ {
		repeated += what
	}
	return repeated
}

func Solve() string {
	a := []byte("")
	s := []byte("527918932189")
	for i := 1; i < len(s)-1; i++ {
		first, _ := strconv.Atoi(string(s[i-1]))
		next, _ := strconv.Atoi(string(s[i+1]))
		desired, _ := strconv.Atoi(string(s[i]))
		if math.Abs(float64(first-next)) == float64(desired) {
			a = append(a, s[i])
		}
	}
	return string(a)
}
