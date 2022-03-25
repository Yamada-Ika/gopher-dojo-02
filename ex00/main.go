package main

import (
	"fmt"
	"os"

	"example.com/ex00/typing"
)

func main() {
	config := typing.NewConfig()
	if err := config.Setquestions("/usr/share/dict/web2"); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	config.SetGameTime(30)
	result, err := config.StartGame()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println("\nTime's up! Score: ", result)
}
