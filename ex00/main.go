package main

import (
	"fmt"
	"os"

	"example.com/ex00/typing"
)

func main() {
	config, err := typing.SetUp("/usr/share/dict/web2")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	// config.SetGameTime(time.Second * 5)
	// config.SetIO(os.Stdin, os.Stdout)
	if err := typing.Start(config); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
