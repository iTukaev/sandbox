package main

import (
	"bufio"
	"errors"
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

type Group struct {
	StudentsByName map[string]Student
}

func NewStudent(name string, age int, grade int) *Student {
	var s = Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
	return &s
}

func consoleInput() (name string, age int, grade int, err error) {
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

func addStudentsToGroup(g *Group) {
	var s Student
	g.StudentsByName = make(map[string]Student)

	for {
		if name, age, grade, err := consoleInput(); err == io.EOF {
			return
		} else {
			s = *NewStudent(name, age, grade)
		}
		if err := g.Put(&s); err != nil {
			fmt.Println(err)
		}
	}
}

func (g *Group) Put(s *Student) error {
	if _, ok := g.StudentsByName[s.Name]; !ok {
		g.StudentsByName[s.Name] = *s
	} else {
		return errors.New("invalid student, repeat enter")
	}
	return nil
}

func (g *Group) GetAll() {
	fmt.Println("Все студенты группы:")
	for key := range g.StudentsByName {
		fmt.Println(g.StudentsByName[key].Name, g.StudentsByName[key].Age, g.StudentsByName[key].Grade)
	}
}

func main() {
	fmt.Println("Введите по очереди: Имя Возраст Курс:")
	fmt.Println("Для выхода нажмите Ctrl+Z и Enter")

	var group = Group{}
	addStudentsToGroup(&group)
	group.GetAll()
}
