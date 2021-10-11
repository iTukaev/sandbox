//go run 26.go -first 123.txt -second 124.txt -result 125.txt

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readFile(name string) string {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}

func main () {
	var flags = []string{"", "", ""}

	flag.StringVar(&flags[0], "first", "", "first file name")
	flag.StringVar(&flags[1], "second", "", "second file name")
	flag.StringVar(&flags[2], "result", "", "result file name")
	flag.Parse()

	var  checkFiles []bool
	var result []string

	for i := range flags {
		checkFiles = append(checkFiles, flags[i] != "")
	}

	if checkFiles[0] {
		result = append(result, readFile(flags[0]))
	}

	if checkFiles[1] {
		result = append(result, readFile(flags[1]))
	}

	resultString := strings.Join(result, "\n")

	if !checkFiles[2] {
		fmt.Println("Результат чтения:")
		fmt.Println(resultString)
	} else {
		file, err := os.Create(flags[2])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		b := bytes.NewBufferString(resultString)

		if err := ioutil.WriteFile(flags[2], b.Bytes(), 0777); err != nil {
			log.Fatal(err)
		}
	}
}
