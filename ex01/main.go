package main

import (
	"context"
	"fmt"
	"os"

	"example.com/ex01/divget"
)

func main() {
	url, divN, err := divget.Parse(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	if divget.ConfigDivN(divN) {
		os.Exit(1)
	}
	if err := divget.Run(context.Background(), url, divN); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
