package cipher

type Shift struct {
	numshift int
}

//func NewShift(num int): Createas a new object of type Cipher with num as it's shift value
func NewShift(num int) Cipher {
	if num == 0 || num < -25 || num > 25 {
		return nil
	}
	var shift Cipher = &Shift{num}
	return shift
}

//func Encode(str string): encodes a string with a shift cipher.numshift
func (cipher Shift) Encode(str string) string {
	var plain []rune = []rune(FormatString(str))
	ciphered := make([]rune, 0)

	for i := 0; i < len(plain); i++ {
		res := rune((int(plain[i])-97+cipher.numshift)%26 + 97)
		if res > 96 {
			ciphered = append(ciphered, res)
		} else {
			res = rune((int(plain[i])-97+cipher.numshift)%26 + 123)
			ciphered = append(ciphered, res)
		}
	}
	return string(ciphered)
}

//func Decode(str string): decodes a string with a shift cipher.numshift
func (cipher Shift) Decode(str string) string {

	var ciphered []rune = []rune(str)
	deciphered := make([]rune, 0)

	for i := 0; i < len(ciphered); i++ {
		r := ciphered[i] - rune(cipher.numshift)
		if r > 96 && r < 123 {
			deciphered = append(deciphered, r)
		} else if r > 122 {
			r = ciphered[i] - rune(26+cipher.numshift)
			deciphered = append(deciphered, r)
		} else if r < 96 {
			r = ciphered[i] + rune(26-cipher.numshift)
			deciphered = append(deciphered, r)
		}
	}
	return string(deciphered)
}
