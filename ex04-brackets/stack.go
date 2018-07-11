package brackets

type Stack struct {
	stackarr []int
}

//func New() *Stack: creates an instance of Stack
func New() *Stack {
	return &Stack{}
}

//func (stack *Stack) Push(num int) : append a num to stackarr array
func (stack *Stack) Push(num int) {
	stack.stackarr = append(stack.stackarr, num)
}

//func (stack *Stack) Pop() int : returns an element that was removed from the end of the stackarr
func (stack *Stack) Pop() int {
	num := stack.stackarr[len(stack.stackarr)-1]
	stack.stackarr = stack.stackarr[:len(stack.stackarr)-1]
	return num
}
