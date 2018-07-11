package cipher

type Vigenere struct {
	str string
}

//func Encode(str string): encodes str by a given string cipher.str

func (cipher Vigenere) Encode(str string) string {
	var plain []rune = []rune(FormatString(str))
	ciphered := make([]rune, 0)

	for i := 0; i < len(plain); i++ {
		shift := NewShift(int(cipher.str[i%len(cipher.str)] - 97))
		var nextstr string
		if shift != nil {
			nextstr = shift.Encode(string(plain[i]))
		} else {
			nextstr = string(plain[i])
		}
		next := []rune(nextstr)
		ciphered = append(ciphered, next[0])
	}
	return string(ciphered)
}

//func Decode(str string): decodes str by a given key cipher.str

func (cipher Vigenere) Decode(str string) string {
	var ciphered []rune = []rune(str)
	decoded := make([]rune, 0)

	for i := 0; i < len(str); i++ {
		var nextstr string
		shift := NewShift(-int(cipher.str[i%len(cipher.str)] - 97))
		if shift != nil {
			nextstr = shift.Encode(string(ciphered[i]))
		} else {
			nextstr = string(ciphered[i])
		}
		next := []rune(nextstr)
		decoded = append(decoded, next[0])
	}
	return string(decoded)

}

//func New Vigenere: creates Vigenere Object of Struct Vigenere

func NewVigenere(str string) Cipher {
	if NonZero(str) == 0 || str != FormatString(str) || len(str) == 0 {
		return nil
	} else {
		var vigenere Cipher = &Vigenere{str}
		return vigenere
	}
}

//func New Vigenere: checks whether NewVigenere argument is not a sequence of a's

func NonZero(str string) int {
	var sum int
	for i := 0; i < len(str); i++ {
		sum += int(str[i])
	}

	if len(str) == 0 || sum/len(str) == 97 {
		return 0
	} else {
		return 1
	}

}
