package main

import (
	"fmt"
	"strings"
)

type StringStack struct {
	data []string
}

func (stack *StringStack) Push(char string) {
	stack.data = append(stack.data, char)
}

func (stack *StringStack) Pop() string {
	var popped string
	popped, stack.data = stack.data[len(stack.data)-1], stack.data[:len(stack.data)-1]
	return popped
}

func (stack StringStack) Read() string {
	return stack.data[len(stack.data)-1]
}

func (stack StringStack) Size() int {
	return len(stack.data)
}

func newStack(sliceType string) StringStack {
	s := StringStack{}
	return s
}

func reverseString(inputString string) string {
	stack := newStack("string")
	var reversedSlice []string = []string{}

	for _, charRune := range inputString {
		stack.Push(string(charRune))
	}

	stackSize := stack.Size()
	for i := 0; i < stackSize; i++ {
		reversedSlice = append(reversedSlice, stack.Pop())
	}

	return strings.Join(reversedSlice, "")
}

func main() {
	// myStack := newStack("string")
	// myStack.Push("a")
	// myStack.Push("b")
	// myStack.Push("c")
	// fmt.Println((myStack.Read()))
	// myStack.Pop()
	// fmt.Println((myStack.Read()))
	fmt.Println(reverseString("abcde"))

}
