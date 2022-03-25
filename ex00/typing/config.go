package typing

import (
	"io"
	"os"
	"strings"
	"time"
)

type Config struct {
	questions []string
	gameTime  time.Duration
	r         io.Reader
	w         io.Writer
}

type config struct {
	questions []string
	gameTime  time.Duration
	r         io.Reader
	w         io.Writer
}

func NewConfig() *Config {
	return &Config{r: os.Stdin, w: os.Stdout}
}

func (cf *Config) Setquestions(path string) error {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return err
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	cf.questions = strings.Split(string(b), "\n")
	return nil
}

func (cf *Config) SetGameTime(seconds int64) {
	cf.gameTime = time.Duration(seconds) * time.Second
}

func isNoConfig(cf *Config) bool {
	if cf == nil || cf.questions == nil {
		return true
	}
	return false
}

func copyConfig(cf *Config) *config {
	return &config{
		questions: cf.questions,
		gameTime:  cf.gameTime,
		r:         cf.r,
		w:         cf.w,
	}
}
