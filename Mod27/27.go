package main

import (
	"errors"
	"fmt"
	"io"
	"visibility/input"
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

func addStudentsToGroup(g *Group) {
	var s Student
	g.StudentsByName = make(map[string]Student)

	for {
		if name, age, grade, err := input.ConsoleInput(); err == io.EOF {
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
