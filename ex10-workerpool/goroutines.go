package goroutines

import (
	"strings"
)

//func Process(in chan string) chan string: Returns a channel with alteredstring
func Process(in chan string) chan string {
	var outString string
	var out chan string = make(chan string)
	go func() {
		inString, ok := <-in
		for ok {
			outString = AlterString(inString)
			out <- outString
			inString, ok = <-in
		}
		defer close(out)
	}()
	return out

}

//func AlterString(inStr string) string : Returns a string with parantheses
func AlterString(inStr string) string {
	b := strings.Builder{}
	b.Grow(len(inStr) + 2)
	b.WriteByte('(')
	b.WriteString(inStr)
	b.WriteByte(')')
	return b.String()
}
