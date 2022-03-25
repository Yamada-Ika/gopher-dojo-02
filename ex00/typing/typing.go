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
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
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

// func makeContext(ctx context.Context, time time.Duration) (context.Context, context.CancelFunc) {
// 	ctx, finish := signal.NotifyContext(context.Background(), syscall.SIGQUIT)
// 	ctx, _ = context.WithTimeout(ctx, time)
// 	return ctx, finish
// }

func (cf *Config) StartGame() error {
	if isNoConfig(cf) {
		return nil
	}
	// ctx, finish := makeContext(context.Background(), cf.gameTime)
	ctxRoot, cancelRoot := context.WithCancel(context.Background())
	defer cancelRoot()
	ctxTime, cancelTime := context.WithTimeout(ctxRoot, cf.gameTime)
	defer cancelTime()
	ctxSIG, cancelSIGINT := signal.NotifyContext(ctxTime, syscall.SIGINT)
	defer cancelSIGINT()
	ctxSIG, cancelSIGQUIT := signal.NotifyContext(ctxSIG, syscall.SIGQUIT)
	defer cancelSIGQUIT()

	// defer finish()
	go func() {
		var solveCount uint64
		for {
			select {
			case <-ctxSIG.Done():
				// cancelTime()
				cancelRoot()
				return
			default:
				changeSeed()
				word := getRandomWord(cf.problems)
				printProblem(cf.w, word)
				ans := solve(cf.r)
				if isEmpty(ans) {
					skip(cf.w)
				} else if checkTheAnswer(word, ans) {
					solveCount++
					ctxTime = sendSolveCount(ctxTime, solveCount)
				}
			}
		}
	}()

	<-ctxTime.Done()
	switch ctxTime.Err() {
	case context.Canceled:
		fmt.Println("Canceled by signal")
	case context.DeadlineExceeded:
		solveCount, ok := getSolveCount(ctxTime)
		if ok {
			fmt.Fprintf(cf.w, "\nTime's up! Score: %d\n", solveCount)
		} else {
			fmt.Fprintf(cf.w, "\nScore can not count...\n")
		}
	}

	return nil
}
