package typing

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

var wordArray []string
var wordArrayLen int

func genRandomWordArray() error {
	f, err := os.Open("/usr/share/dict/web2")
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	wordArray = strings.Split(string(b), "\n")
	wordArrayLen = len(wordArray)
	return nil
}

var solveCount int

func typing(ctx context.Context) {
	rand.Seed(time.Now().UnixNano())
	for {
		word := wordArray[rand.Intn(wordArrayLen)]
		fmt.Println(word)
		fmt.Printf("-> ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() != word {
				fmt.Printf("\033[1A\033[3C\033[K")
				continue
			}
			solveCount++
			break
		}
	}
}

func StartGame() {
	genRandomWordArray()
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()
	go typing(ctx)
	select {
	case <-ctx.Done():
		fmt.Printf("\nTime's up! Score: %d\n", solveCount)
	}
}
