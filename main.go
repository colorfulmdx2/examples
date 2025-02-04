package main

import "fmt"

func main() {
	err := do()
	if err != nil {
		fmt.Println("ERROR")
	} else {
		fmt.Println("NO ERROR")
	}
}

func do() error {
	var p *MyError
	if false {
		p = &MyError{"error"}
	}

	return p
}

type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return e.msg
}
