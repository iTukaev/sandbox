package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(path, printFiles) //
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(path string, printFiles bool) error { //in *os.File,
	catalogueContent, err := os.ReadDir(path)
	if err != nil {
		return nil
	}
	for num, name := range catalogueContent {

		if printFiles || name.IsDir() {
			if num != len(catalogueContent)-1 {
				countOfSeparator := strings.Count(path, string(os.PathSeparator))
				tabSwitch := ""
				for i := 0; i < countOfSeparator; i++ {
					tabSwitch += "│\t"
				}
				fmt.Print(tabSwitch)
			}
			if num != len(catalogueContent)-1 {
				fmt.Print("├───")
			} else {
				fmt.Print("└───")
			}
			fmt.Println(name.Name())
			err := dirTree(path+string(os.PathSeparator)+name.Name(), printFiles)
			if err != nil {
				return nil
			}
		}
	}
	return nil
}