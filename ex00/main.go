package main

import (
	"fmt"
	"os"

	"example.com/ex00/typing"
)

func main() {
	if err := typing.StartGame(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
