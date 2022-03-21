package main

import (
	"fmt"
	"os"

	"example.com/ex00/typing"
)

// NewConfig() コンストラクタ
// NewDefaultConfig() 別のコンストラクタ
// (cf *typing.Config) SetProblems(filePath) error 英単語をセットするメソッド
// (cf *typing.Config) SetGameTime(seconds) ゲーム時間をセットするメソッド
// (cf *typing.Config) Config() *typing.config //nilチェックするメソッド
// (cf *typing.config) StartGame() error ゲーム時間を開始するメソッド
func main() {
	config := typing.NewConfig()
	if err := config.SetProblems("/usr/share/dict/web2"); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	config.SetGameTime(30)
	if err := config.StartGame(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
