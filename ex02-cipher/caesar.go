package cipher

type Caesar struct {
}

//func Encode(str string): encodes the string with shift 3

func (cipher Caesar) Encode(str string) string {
	c := NewShift(3)
	ciphered := c.Encode(str)
	return ciphered
}

//func Decode(str string): decodes the string with shift 3

func (cipher Caesar) Decode(str string) string {
	c := NewShift(3)
	deciphered := c.Decode(str)
	return deciphered
}

//func NewCaesar(): creates an Object of type Cipher

func NewCaesar() Cipher {
	var caesar Cipher = &Caesar{}
	return caesar
}
