package simple_text_editor

import (
	"fmt"
	"strconv"

	"github.com/golang-collections/collections/stack"
)

func SimpleTextEditorStack(input string, commands []string) []string {
	returnArray := []string{}
	s := input
	history := stack.New()
	history.Push(input)
	for _, v := range commands {
		switch v[0] {
		case '1': // append
			stack_add(&s, v)
			history.Push(s)
		case '2': // delete
			stack_delete(&s, v)
			history.Push(s)
		case '3': // print
			result := stack_print(&s, v)
			returnArray = append(returnArray, result)
		case '4': // undo
			stack_undo(&s, history)
		}
	}
	return returnArray
}
func stack_add(s *string, append string) {
	*s = fmt.Sprintf("%s%s", *s, append[2:])
}
func stack_delete(s *string, delete string) {
	index, err := strconv.Atoi(delete[2:])
	if err != nil {
		fmt.Println(err)
	}

	if index > len(*s) {
		*s = ""
		return
	}

	*s = (*s)[0 : len(*s)-index]
}
func stack_print(s *string, print string) string {
	index, err := strconv.Atoi(print[2:])
	if err != nil {
		fmt.Println(err)
	}
	return string((*s)[index-1])
}
func stack_undo(s *string, history *stack.Stack) {
	history.Pop()
	top := history.Peek()
	result, ok := top.(string)
	if !ok {
		*s = ""
		return
	}
	*s = result
}
