package main

import (
	"fmt"
	"strings"
)

func Stats(values []int) (min int, max int, average float64, ok bool) {
	if len(values) == 0 {
		return 0, 0, 0, false
	}
	min = values[0]
	max = values[0]
	sum := 0
	for _, v := range values {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		sum += v
	}
	average = float64(sum) / float64(len(values))
	ok = true
	return min, max, average, ok
}

func GradeCounts(values []int) map[int]int {
	counts := make(map[int]int)
	for _, v := range values {
		counts[v]++
	}
	return counts
}

func WordCounts(text string) map[string]int {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}
	return counts
}

func main() {
	marks := [5]int{5, 4, 3, 5, 4}
	marksSlice := append(marks[:], 2, 5)

	fmt.Printf("Исходный массив: %v\n", marks)
	fmt.Printf("Срез после append: %v\n", marksSlice)
	fmt.Printf("Длина среза: %d, ёмкость: %d\n", len(marksSlice), cap(marksSlice))

	min, max, avg, ok := Stats(marksSlice)
	fmt.Printf("Stats (оценки): min = %d, max = %d, average = %v, ok = %t\n", min, max, avg, ok)

	emptyMin, emptyMax, emptyAvg, emptyOk := Stats([]int{})
	fmt.Printf("Stats (пустой срез): min = %d, max = %d, average = %v, ok = %t\n", emptyMin, emptyMax, emptyAvg, emptyOk)

	grades := GradeCounts(marksSlice)
	for grade := 2; grade <= 5; grade++ {
		fmt.Printf("%d: %d\n", grade, grades[grade])
	}

	demoMap := make(map[string]int)
	demoMap["item"] = 10
	demoMap["item"] = 25
	val, exists := demoMap["item"]
	fmt.Printf("До удаления: значение = %d, ok = %t\n", val, exists)
	delete(demoMap, "item")
	val, exists = demoMap["item"]
	fmt.Printf("После удаления: значение = %d, ok = %t\n", val, exists)

	text := "Go go язык Go язык"
	words := WordCounts(text)
	fmt.Printf("go: %d; язык: %d\n", words["go"], words["язык"])

	greeting := "Привет"
	fmt.Printf("len(\"%s\") = %d байтов, len([]rune(\"%s\")) = %d рун\n", greeting, len(greeting), greeting, len([]rune(greeting)))
}
