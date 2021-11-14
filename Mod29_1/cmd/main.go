package main

import (
	"fmt"
	"sandbox/Mod29_1/pkg/steps"
)

func main() {
	fmt.Println("Input numbers or \"stop\" to break function")
	for {
		fc, err := steps.WaitNum()
		if err != nil {
			break
		}
		sc := steps.Quad(fc)
		tc := steps.Multiplication(sc)
		fmt.Println("Произведение:", <-tc)
	}
}