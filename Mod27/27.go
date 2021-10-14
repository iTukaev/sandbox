package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name  string
	age   int
	grade int
}

func (s *Student) put() error {
	var err error = nil
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		txt := sc.Text()
		words := strings.Fields(txt)
		s.name = words[0]
		s.age, _ = strconv.Atoi(words[1])
		s.grade, _ = strconv.Atoi(words[2])
	} else {
		err = io.EOF
	}
	return err
}

func (s *Student) get() {
	fmt.Printf("Имя: %s. Возраст: %d. Курс: %d\n", s.name, s.age, s.grade)
}

func newStudents() (allStudents map[int]Student) {
	allStudents = make(map[int]Student)
	var s Student
	index := 1
	for {
		if err := s.put(); err != io.EOF {
			allStudents[index] = s
			index++
		} else {
			break
		}
	}
	return allStudents
}

func indexSort(m *map[int]Student) (index []int) {
	index = make([]int, len(*m))

	for key := range *m {
		index[key-1] = key
	}
	return
}

func main() {
	fmt.Println("Введите по очереди: Имя Возраст Курс:")
	fmt.Println("Для выход анажмите Ctrl+Z и Enter")

	var allStudents map[int]Student
	allStudents = make(map[int]Student)

	allStudents = newStudents()

	keys := indexSort(&allStudents)

	for _, key := range keys {
		fmt.Print("Студент №", key, "\t")
		s := allStudents[key]
		s.get()
	}
}
