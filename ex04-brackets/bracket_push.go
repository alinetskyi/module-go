package brackets

//func Bracket(str string) (bool, error) : return true if all brackets are closed, otherwise false
func Bracket(str string) (bool, error) {
	var stack *Stack = New()
	if len(str) == 0 {
		return true, nil
	}
	for r := range str {
		switch string(str[r]) {
		case "{":
			stack.Push(1)
		case "[":
			stack.Push(2)
		case "(":
			stack.Push(3)
		case "}":
			if len(stack.stackarr) != 0 {
				tmp := stack.Pop()
				if tmp != 1 {
					return false, nil
				}
			}
		case "]":
			if len(stack.stackarr) != 0 {
				tmp := stack.Pop()
				if tmp != 2 {
					return false, nil
				}
			}
		case ")":
			if len(stack.stackarr) != 0 {
				tmp := stack.Pop()
				if tmp != 3 {
					return false, nil
				}
			}
		default:
			return false, nil
		}
	}
	if len(stack.stackarr) != 0 {
		return false, nil
	}
	return true, nil
}
