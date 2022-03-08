package typing

import (
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Config struct {
	asset    []string
	gameTime time.Duration
	r        io.Reader
	w        io.Writer
}

func SetUp(filePath string) (cf *Config, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return &Config{
		asset:    strings.Split(string(b), "\n"),
		gameTime: time.Second * 30,
		r:        os.Stdin,
		w:        os.Stdout,
	}, nil
}

func (cf *Config) SetGameTime(time time.Duration) {
	cf.gameTime = time
}

func (cf *Config) SetIO(r io.Reader, w io.Writer) {
	cf.r = r
	cf.w = w
}

// func GetWordAsset(filePath string) (words []string, err error) {
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
