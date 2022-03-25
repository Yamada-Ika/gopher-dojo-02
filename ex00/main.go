package main

import (
	"fmt"
	"os"
	"runtime"

	"example.com/ex00/typing"
)

func checkGoroutineLeak() {
	fmt.Println(runtime.NumGoroutine())
}

func main() {
	checkGoroutineLeak()
	defer checkGoroutineLeak()
	config := typing.NewConfig()
	if err := config.SetProblems("/usr/share/dict/web2"); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	config.SetGameTime(30)
	if err := config.StartGame(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
