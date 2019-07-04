package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// サイズが1よりも大きいチャネルを作成
	signals := make(chan os.Signal, 1)
	// SIGINT(Ctrl+C) を受け取る
	signal.Notify(signals, syscall.SIGINT)

	// シグナルが来るまで待つ
	fmt.Println("Waiting SIGINT (CTRL+C)")
	<-signals
	fmt.Println("SIGINT arrived")
}
