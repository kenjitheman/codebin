package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	text, err := print_some_value()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(text)
}

func print_some_value() (string, error) {
	return "some value", nil
}
