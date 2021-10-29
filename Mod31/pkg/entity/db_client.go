package entity

type Department struct {
	Users map[int]*User
}

func NewDepartment() *Department {
	dep := new(Department)
	dep.Users = make(map[int]*User)
	return dep
}