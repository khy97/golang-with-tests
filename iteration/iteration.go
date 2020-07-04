package iteration

// Repeat -> Repeats the number 5 times
func Repeat(c string, count int) (repeated string) {
	for i := 0; i < count; i++ {
		repeated += c
	}
	return
}
