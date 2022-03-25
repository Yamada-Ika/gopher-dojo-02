package typing

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func checkTheAnswer(exp, ans string) bool {
	return exp == ans
}

func changeSeed() {
	rand.Seed(time.Now().UnixNano())
}

func makeAnswerCh(r io.Reader) <-chan string {
	ch := make(chan string, 1)

	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}()

	return ch
}

func (cf *config) getQuestion() string {
	changeSeed()
	return cf.questions[rand.Intn(len(cf.questions))]
}

func (cf *config) printQuestion(question string) {
	fmt.Fprintf(cf.w, "%s\n->", question)
}

func (c *Config) StartGame() (result uint64, err error) {
	if isNoConfig(c) {
		return 0, configErr
	}

	cf := copyConfig(c)

	answerCh := makeAnswerCh(cf.r)
	done := make(chan bool)

	go func() {
		for {
			q := cf.getQuestion()
			cf.printQuestion(q)
			select {
			case <-done:
				return
			case ans := <-answerCh:
				if checkTheAnswer(q, ans) {
					result++
				}
			}
		}
	}()

	<-time.After(cf.gameTime)
	done <- true
	time.Sleep(time.Second)

	return result, nil
}
