package typing

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os/signal"
	"syscall"
	"time"
)

func printProblem(w io.Writer, problem string) {
	fmt.Fprintf(w, "%s\n->", problem)
}

func solve(r io.Reader) (ans string) {
	scanner := bufio.NewScanner(r)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func checkTheAnswer(exp, ans string) bool {
	return exp == ans
}

func isEmpty(s string) bool {
	return s == ""
}

func getSolveCount(ctx context.Context) (uint64, bool) {
	val, ok := ctx.Value("solveCount").(uint64)
	return val, ok
}

func skip(w io.Writer) {
	fmt.Fprintf(w, "\n")
}

func changeSeed() {
	rand.Seed(time.Now().UnixNano())
}

func Start(cf *Config) error {
	ctx, finish := signal.NotifyContext(context.Background(), syscall.SIGQUIT)
	defer finish()

	ctx, _ = context.WithTimeout(ctx, cf.gameTime)

	go func() {
		var count uint64
		changeSeed()
		for {
			word := getRandomWord(cf.asset)
			printProblem(cf.w, word)
			ans := solve(cf.r)
			if isEmpty(ans) {
				skip(cf.w)
				continue
			}
			if checkTheAnswer(word, ans) {
				count++
				ctx = context.WithValue(ctx, "solveCount", count)
			}
		}
	}()

	select {
	case <-ctx.Done():
		solveCount, ok := getSolveCount(ctx)
		if ok {
			fmt.Fprintf(cf.w, "\nTime's up! Score: %d\n", solveCount)
		} else {
			fmt.Fprintf(cf.w, "\nScore can not count...\n")
		}
	}
	return nil
}
