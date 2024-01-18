package iteration

func Repeat(character string, repeatCount int) string {
	// this is a bare assignment ig
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
