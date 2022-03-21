package typing

import (
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Config struct {
	problems []string
	gameTime time.Duration
	r        io.Reader
	w        io.Writer
}

type config struct {
	problems []string
	gameTime time.Duration
	r        io.Reader
	w        io.Writer
}

func NewConfig() *Config {
	return &Config{r: os.Stdin, w: os.Stdout}
}

func (cf *Config) SetProblems(path string) error {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return err
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	cf.problems = strings.Split(string(b), "\n")
	return nil
}

func (cf *Config) SetGameTime(seconds int64) {
	cf.gameTime = time.Duration(seconds) * time.Second
}

// func SetUp(filePath string) (*Config, error) {
// 	f, err := os.Open(filePath)
// 	defer f.Close()
// 	if err != nil {
// 		return nil, err
// 	}
// 	b, err := ioutil.ReadAll(f)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Config{
// 		problems: strings.Split(string(b), "\n"),
// 		gameTime: time.Second * 30,
// 		r:        os.Stdin,
// 		w:        os.Stdout}, nil
// }

// func (cf *Config) SetGameTime(time time.Duration) {
// 	cf.gameTime = time
// }

// func (cf *Config) SetIO(r io.Reader, w io.Writer) {
// 	cf.r = r
// 	cf.w = w
// }

// func GetWordproblems(filePath string) (words []string, err error) {
// 	f, err := os.Open(filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	b, err := ioutil.ReadAll(f)
// 	if err != nil {
// 		return nil, err
// 	}
// 	wordList := strings.Split(string(b), "\n")
// 	defer f.Close()
// 	return wordList, nil
// }

func getRandomWord(words []string) string {
	return words[rand.Intn(len(words))]
}
