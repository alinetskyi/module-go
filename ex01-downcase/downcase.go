package downcase

/*
func main() {
	fmt.Println(Downcase("VXZEWASDKdwewq321321+ewq344@#@321FXDCVX"))
}
*/

func Downcase(str string) (string, error) {
	var resstr []rune = []rune(str)

	for r := range str {
		if resstr[r] > 64 && resstr[r] < 91 {
			resstr[r] += 32
		}
	}
	return string(resstr), nil

}
