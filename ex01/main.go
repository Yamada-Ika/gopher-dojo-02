package main

import (
	"fmt"
	"os"

	"example.com/ex01/divget"
)

func main() {
	if err := divget.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
}
