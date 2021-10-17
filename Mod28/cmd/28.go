package main

import (
	"fmt"
	group "sandbox/Mod28/pkg/group"
)

func main() {
	fmt.Println("Введите по очереди: Имя Возраст Курс:")
	fmt.Println("Для выхода нажмите Ctrl+Z и Enter")

	gr := group.NewGroup()
	group.AddAllStudentsToGroup(&gr)
	group.PrintAllStudentsOfGroup(&gr)
}
