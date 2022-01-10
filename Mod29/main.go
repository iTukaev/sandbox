package main

import (
	"fmt"
	"io"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Input number or \"stop\" to break function")
	for {
		num, err := NumInsert()
		if err == io.EOF {
			break
		}

		fc := make(chan int)
		sc := make(chan int)
		tc := make(chan int)

		go Quad(fc, sc)
		go Multiplication(sc, tc)

		fc <- num
		fmt.Println("Multiple:", <-tc)
	}
}

func NumInsert() (int, error) {
	for {
		fmt.Print("Input: ")
		var numTxt string
		if _, err := fmt.Scan(&numTxt); err != nil {
			fmt.Println("Input number or \"stop\" to break function")
			continue
		}

		if numTxt == "stop" {
			return 0, io.EOF
		}

		num, err := strconv.Atoi(numTxt)
		if err != nil {
			fmt.Println("Input correct number or \"stop\" to break function")
			continue
		}
		return num, nil
	}
}

func Quad(in, out chan int) {
	num := int(math.Pow(float64(<-in), 2.0))
	fmt.Println("Quad:", num)

	out <- num
}

func Multiplication(in, out chan int){
	num := (<-in) * 2
	fmt.Println("Multiple:", num)

	out <- num
}