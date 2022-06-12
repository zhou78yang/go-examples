package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 通过channel接收Signal
	sigs := make(chan os.Signal, 1)
	// 注册给定的通道，用于接收特定信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
