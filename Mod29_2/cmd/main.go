package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main()  {

	exitWait()

	fmt.Println("Пока!")
}

func exitWait() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		os.Interrupt,
		syscall.SIGINT)

	i := 1

	for i < 30000 {
		select {
		case <- sigChan:
			fmt.Println("Выхожу, ожидайте завершения.")
			time.Sleep(time.Second)
			return
		default:
			fmt.Println(<- square(i))
		}
		i++
	}
}

func square(num int) <-chan int {
	chQuad := make(chan int)

	go func() {
		chQuad <- num * num
	}()
	return chQuad
}