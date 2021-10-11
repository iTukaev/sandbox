//go run 26.go -first 123.txt -second 124.txt -result 125.txt

package main

import (
	"flag"
	"fmt"
)

func main () {
	var flag1, flag2, flag3 string

	//flag.StringVar(&flag1, "", "","")
	//flag.StringVar(&flag2, "", "","")
	//flag.StringVar(&flag3, "", "","")

	//flag.Parse()
	flag.StringVar(&flag1, "first", "", "first file name")
	flag.StringVar(&flag2, "second", "", "second file name")
	flag.StringVar(&flag3, "result", "", "result file name")

	flag.Parse()
	fmt.Printf("%v\t%v\t%v\t", flag1, flag2, flag3)
}
