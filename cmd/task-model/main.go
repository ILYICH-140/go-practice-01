package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

func (t Task) Summary() string {
	return fmt.Sprintf("#%d %s [%s], приоритет %d", t.ID, t.Title, t.Status, t.Priority)
}

func (t *Task) Complete() {
	t.Status = "done"
}

func (t *Task) SetPriority(priority int) bool {
	if priority < 1 || priority > 3 {
		return false
	}
	t.Priority = priority
	return true
}

func findTaskByID(tasks []Task, id int) *Task {
	for i, task := range tasks {
		if task.ID == id {
			return &tasks[i]
		}
	}
	return nil
}

func tasksByStatus(tasks []Task, status string) []Task {
	arr := []Task{}
	for _, task := range tasks {
		if task.Status == status {
			arr = append(arr, task)
		}
	}
	return arr
}

func main() {
	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}

	task2 := findTaskByID(tasks, 2)
	task2.Complete()
	task2.SetPriority(3)

	for _, task := range tasks {
		fmt.Println(task.Summary())
	}
	// var count int = 0
	// for _, task := range tasks {
	// 	if task.Status == "done" {
	// 		count++
	// 	}
	// }
	// fmt.Println("Количество завершённых задач: ", count)
	fmt.Println("Количество завершённых задач: ", len(tasksByStatus(tasks, "done")))
}
