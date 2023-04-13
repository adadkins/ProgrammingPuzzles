package hello_worldinator

import "fmt"

func HelloWorldinator() func() {
	f := func() {
		fmt.Println("Hello World!")
	}
	return f
}
