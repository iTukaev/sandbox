package group

import (
	"fmt"
	"io"
	student "sandbox/Mod28/pkg/Student"
)

type Group map[string]*student.Student

func NewGroup() Group {
	return make(map[string]*student.Student, 0)
}

func (g Group) Put(s *student.Student) error {
	if _, ok := g[s.Name]; !ok {
		g[s.Name] = s
	} else {
		return fmt.Errorf("invalid student, repeat enter")
	}
	return nil
}

func (g Group) Get(name string) (*student.Student, error){
	s, ok := g[name]
	if !ok {
		return nil, fmt.Errorf("no suth user")
	}
	return s, nil
}

func PrintAllStudentsOfGroup(g *Group) {
	fmt.Println("Все студенты группы:")
	for _, st := range *g {
		fmt.Println(st.Name, st.Age, st.Grade)
	}
}

func AddAllStudentsToGroup(g *Group) {
	var err error = nil

	for {
		s := student.NewStudent()
		err = student.ConsoleInput(s)
		if err == io.EOF {
			return
		}
		err = g.Put(s)
		if err != nil {
			fmt.Println(err)
		}
	}
}

