package main

import "fmt"

func epid() {
	const language = "Go"
	var name string = "Илья"
	completed := 3
	var ready bool = true
	var hours float64 = 4.2

	var str string = "Строки"
	var num int = 42
	var bl bool = true
	var div float64 = 1.95629485184
	var div3 float64 = 141.0 / 71
	tip := false
	stri := "fine"

	fmt.Println("Изучаем", language)
	fmt.Printf("Участник: %s\n", name)
	fmt.Printf("Выполнено: %d\n", completed)
	fmt.Printf("Готов продолжать: %t\n", ready)
	fmt.Printf("Продолжительность: %.1f часа\n", hours)

	completed = 2
	fmt.Printf("После изменения: %d\n", completed)
	fmt.Print("==============================\n")
	fmt.Printf("Эксперимент по выводу: %s\n", str)
	fmt.Printf("Целого числа: %d\n", num)
	fmt.Printf("Логического значения: %t\n", bl)
	fmt.Printf("Дробное число: %g\n", div)
	fmt.Printf("Дробное число с 3 знаками после точки: %.3f \n", div3)
	fmt.Printf("Тип значения: %T\n", tip)
	fmt.Printf("Строку в кавычках: %q\n", stri)
	fmt.Print("==============================\n")
	exp()
}
