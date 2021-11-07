package entity

import "sandbox/Mod30/pkg/groupServise"

type Department struct {
	Users map[int]*groupServise.User
}

func NewDepartment() (dep *Department) {
	dep = new(Department)
	dep.Users = make(map[int]*groupServise.User)
	return dep
}