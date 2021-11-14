package student

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

func NewStudent() *Student {
	return &Student{}
}

func ConsoleInput(student *Student) (err error) {
	sc := bufio.NewScanner(os.Stdin)
Beginning:
	if sc.Scan() {
		txt := sc.Text()
		words := strings.Fields(txt)
		if len(words) == 3 {
			student.Name = words[0]
			student.Age, _ = strconv.Atoi(words[1])
			student.Grade, _ = strconv.Atoi(words[2])
		} else {
			fmt.Println("Input error, repeat")
			goto Beginning
		}
	} else {
		err = io.EOF
	}
	return
}
