package main

import "fmt"

func main() {
	var totalMinutes int
	fmt.Print("Введите количество минут: ")

	// Этот блок ввода дан готовым; вычисления добавьте самостоятельно.
	_, err := fmt.Scan(&totalMinutes)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if totalMinutes < 0 {
		fmt.Println("Ошибка: количество минут не должно быть отрицательным.")
		return
	}

	var minutesInHour int = 60
	var fullHours int = totalMinutes / minutesInHour
	var remainingMinutes int = totalMinutes - 60*fullHours
	// TODO: объявите константу — количество минут в одном часе.
	// TODO: вычислите полные часы и оставшиеся минуты.
	// TODO: выведите результаты в формате задания 3.
	fmt.Printf("Полных часов: %d. Оставшихся минут: %d.\n", fullHours, remainingMinutes)
}
