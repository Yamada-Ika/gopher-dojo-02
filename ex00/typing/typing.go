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

func isNoConfig(cf *Config) bool {
	if cf == nil || cf.problems == nil {
		return true
	}
	return false
}

func sendSolveCount(ctx context.Context, count uint64) context.Context {
	return context.WithValue(ctx, "solveCount", count)
}

func waitGame(ctx context.Context) {
	<-ctx.Done()
}

func makeContext(ctx context.Context, time time.Duration) (context.Context, context.CancelFunc) {
	ctx, finish := signal.NotifyContext(context.Background(), syscall.SIGQUIT)
	ctx, _ = context.WithTimeout(ctx, time)
	return ctx, finish
}

func (cf *Config) StartGame() error {
	if isNoConfig(cf) {
		return nil
	}
	ctx, finish := makeContext(context.Background(), cf.gameTime)
	defer finish()
	go func() {
		var solveCount uint64
		changeSeed()
		for {
			word := getRandomWord(cf.problems)
			printProblem(cf.w, word)
			ans := solve(cf.r)
			if isEmpty(ans) {
				skip(cf.w)
				continue
			}
			if checkTheAnswer(word, ans) {
				solveCount++
				ctx = sendSolveCount(ctx, solveCount)
			}
		}
	}()
	waitGame(ctx)
	solveCount, ok := getSolveCount(ctx)
	if ok {
		fmt.Fprintf(cf.w, "\nTime's up! Score: %d\n", solveCount)
	} else {
		fmt.Fprintf(cf.w, "\nScore can not count...\n")
	}
	return nil
}
