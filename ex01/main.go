package main

import (
	"fmt"
	"os"

	"example.com/ex01/divget"
)

func main() {
	url, err := divget.GetURL(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	var parallelN uint64
	parallelN, err = divget.GetParallelN(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	if err := divget.Run(url, parallelN); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
