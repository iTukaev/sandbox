//go run 26.go 123.txt 124.txt 125.txt

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readFiles(files []string) string {
	readStrings := make([]string, len(files), len(files))

	for i, name := range files {
		if i < 2 {
			file, err := os.Open(name)
			if err != nil {
				log.Fatal(err)
			}

			b, err := ioutil.ReadAll(file)
			if err != nil {
				log.Fatal(err)
			}
			readStrings[i] = string(b)
			file.Close()
		}
	}

	return stringsConcatenate(readStrings)
}

func stringsConcatenate(readStrings []string) string {
	return strings.Join(readStrings, "\n")
}

func newFileCreate(files []string) error {
	file, err := os.Create(files[2])
	if err != nil {
		return err
	}
	defer file.Close()

	b := bytes.NewBufferString(readFiles(files))

	if err := ioutil.WriteFile(files[2], b.Bytes(), 0777); err != nil {
		return err
	}
	return nil
}

func main() {
	filesNames := os.Args[1:]

	switch len(filesNames) {
	case 0:
		fmt.Println("А где же список файлов?")
	case 1, 2:
		fmt.Println("Содержимое файла(ов):\n\n", readFiles(filesNames))
	case 3:
		if err := newFileCreate(filesNames); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Результат выполнения в файле:", filesNames[2])
	}
}
