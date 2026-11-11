package main

import "fmt"

func main() {
	var name string = "Илья"
	var group string = "Учебная-01"
	var course string = "Бэкенд-разработка на языке Go"
	completed := 1
	var interest bool = true

	fmt.Println("Участник: ", name)
	fmt.Printf("Группа: %s\n", group)
	fmt.Printf("Курс: %s\n", course)
	fmt.Printf("Завершено работ: %d\n", completed)
	fmt.Printf("Интерес к бэкенду: %t\n", interest)
}
