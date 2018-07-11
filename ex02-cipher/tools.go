package cipher

//func FormatString(str string) : returns formated string with only lowercase letters
func FormatString(str string) string {
	var resstr []rune = []rune(str)
	chars := make([]rune, 0)

	for i := 0; i < len(str); i++ {
		if resstr[i] > 64 && resstr[i] < 91 {
			resstr[i] += 32
			chars = append(chars, resstr[i])
		} else if resstr[i] > 96 && resstr[i] < 123 {
			chars = append(chars, resstr[i])
		} else {

		}
	}
	return string(chars)
}
