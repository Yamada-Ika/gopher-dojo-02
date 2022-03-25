package typing_test

import (
	"testing"

	"example.com/ex00/typing"
	"go.uber.org/goleak"
)

func TestStartGame(t *testing.T) {
	defer goleak.VerifyNone(t)
	cf := typing.NewConfig()
	cf.SetGameTime(5)
	cf.Setquestions("../testdata/test.txt")
	t.Run("normal", func(t *testing.T) {
		cf.StartGame()
	})
}
