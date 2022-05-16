package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	for _, w := range input {
		output = string(w) + output
	}
	return
}
