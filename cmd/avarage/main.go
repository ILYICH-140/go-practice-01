package main

import "fmt"

func main() {
	var num1, num2, num3 int
	var err error
	fmt.Print("Введите три целых числа через пробелл: ")

	// Этот блок ввода дан готовым; вычисления добавьте самостоятельно.
	_, err = fmt.Scan(&num1, &num2, &num3)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целые числа.")
		return
	}

	var summ int = num1 + num2 + num3
	average := float64(summ) / 3

	fmt.Printf("Сумма %d, cреднее %.2f\n", summ, average)
}
