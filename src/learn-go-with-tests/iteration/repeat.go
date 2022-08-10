package iteration

func Repeat(character string, counts int) string {
	var repeated string
	for i := 0; i < counts; i++ {
		repeated += character
	}
	return repeated
}
