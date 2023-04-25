package simple_text_editor

import (
	"fmt"
	"strconv"
)

func SimpleTextEditorArray(input string, commands []string) []string {
	returnArray := []string{}
	s := input
	history := []string{s}
	for _, v := range commands {
		switch v[0] {
		case '1': // append
			array_add(&s, v)
			history = append(history, s)
		case '2': // delete
			array_delete(&s, v)
			history = append(history, s)
		case '3': // print
			result := array_print(&s, v)
			returnArray = append(returnArray, result)
		case '4': // undo
			array_undo(&s, &history)
		}
	}
	return returnArray
}

func array_add(s *string, append string) {
	*s = fmt.Sprintf("%s%s", *s, append[2:])
}
func array_delete(s *string, delete string) {
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
func array_print(s *string, print string) string {
	index, err := strconv.Atoi(print[2:])
	if err != nil {
		fmt.Println(err)
	}
	return string((*s)[index-1])
}
func array_undo(s *string, history *[]string) {
	if len(*history) == 0 {
		return
	}
	*history = (*history)[:len(*history)-1]
	*s = (*history)[len(*history)-1]
}
