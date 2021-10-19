package steps

import (
	"fmt"
	"io"
	"math"
	"strconv"
)

func WaitNum() (chan int, error) {
	firstStepChan := make(chan int)
Begin:
	fmt.Print("Ввод: ")
	var numTxt string
	fmt.Scan(&numTxt)
	if numTxt == "stop" {
		return firstStepChan, io.EOF
	}
	num, err := strconv.Atoi(numTxt)
	if err != nil {
		goto Begin
	}
	go func() {
		firstStepChan <- num
	}()
	return firstStepChan, nil
}

func Quad(firstStepChan chan int) chan int{
	secondStepChan := make(chan int)
	num := int(math.Pow(float64(<-firstStepChan), 2.0))
	go func() {
		secondStepChan <- num
		secondStepChan <- num
	}()
	return secondStepChan
}

func Multiplication(secondStepChan chan int) chan int{
	fmt.Println("Квадрат:", <- secondStepChan)

	thirdStepChan := make(chan int)
	num := (<-secondStepChan)*2
	go func() {
		thirdStepChan <- num
	}()
	return thirdStepChan
}
