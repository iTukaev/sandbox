package input

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ConsoleInput() (name string, age int, grade int, err error) {
	sc := bufio.NewScanner(os.Stdin)
Beginning:
	if sc.Scan() {
		txt := sc.Text()
		words := strings.Fields(txt)
		if len(words) == 3 {
			name = words[0]
			age, _ = strconv.Atoi(words[1])
			grade, _ = strconv.Atoi(words[2])
		} else {
			fmt.Println("Input error, repeat")
			goto Beginning
		}
	} else {
		err = io.EOF
	}
	return
}
