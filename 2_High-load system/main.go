package main

import (
	"bufio"
	"fmt"
	"os"
)

func do(b string) string {
	// Процесс должен начинаться с M и заканчиваться на D
	if len(b) == 0 || b[0] != 'M' || b[len(b)-1] != 'D' {
		return "NO"
	}

	// Текущее состояние задачи
	state := "NONE" // начальное состояние
	lastAction := ""

	for i := 0; i < len(b); i++ {
		action := b[i]
		switch action {
		case 'M': // Запуск задачи
			if state == "NONE" || state == "CANCELLED" || state == "DONE" {
				state = "RUNNING"
			} else {
				return "NO"
			}
		case 'R': // Рестарт задачи
			if state == "RUNNING" {
				state = "RESTARTED"
			} else {
				return "NO"
			}
		case 'C': // Отмена задачи
			if state == "RUNNING" || state == "RESTARTED" {
				state = "CANCELLED"
			} else {
				return "NO"
			}
		case 'D': // Завершение задачи
			if state == "RUNNING" || state == "CANCELLED" {
				state = "DONE"
			} else {
				return "NO"
			}
		default:
			return "NO"
		}

		// Проверка на последовательность действий
		if lastAction == "R" && action != 'C' {
			return "NO"
		}

		// Обновление последнего действия
		lastAction = string(action)
	}

	// Проверка на дублирование действий
	for i := 1; i < len(b); i++ {
		if b[i] == b[i-1] {
			return "NO"
		}
	}

	if state == "DONE" {
		return "YES"
	}
	return "NO"
}

func main() {
	var in = bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscan(in, &t)

	results := make([]string, t)

	for i := 0; i < t; i++ {
		var b string
		fmt.Fscan(in, &b)
		results[i] = do(b)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
