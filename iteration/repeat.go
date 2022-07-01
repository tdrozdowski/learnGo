package iteration

func Repeat(what string, howMany int) string {
	var repeated string
	for i := 0; i < howMany; i++ {
		repeated += what
	}
	return repeated
}
