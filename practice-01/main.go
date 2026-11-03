package main

import "fmt"

func main() {
	const courseName = "Бэкенд-разработка на языке Go"
	var name string = "Роман"
	var age int = 20
	average := 4.75
	active := true

	fmt.Printf("courseName: %v (%T)\n", courseName, courseName)
	fmt.Printf("name: %v (%T)\n", name, name)
	fmt.Printf("age: %v (%T)\n", age, age)
	fmt.Printf("average: %v (%T)\n", average, average)
	fmt.Printf("active: %v (%T)\n", active, active)

	a := 5
	b := 2
	operation := "/"

	switch operation {
	case "+":
		fmt.Printf("Результат: %d\n", a+b)
	case "-":
		fmt.Printf("Результат: %d\n", a-b)
	case "*":
		fmt.Printf("Результат: %d\n", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль")
		} else {
			fmt.Printf("Результат: %v\n", float64(a)/float64(b))
		}
	case "%":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль")
		} else {
			fmt.Printf("Результат: %d\n", a%b)
		}
	default:
		fmt.Printf("Ошибка: неизвестная операция %q\n", operation)
	}

	n := 10

	if n < 1 || n > 100 {
		fmt.Printf("Ошибка: число n = %d находится вне диапазона от 1 до 100\n", n)
	} else {
		sum := 0
		evenCount := 0
		for i := 1; i <= n; i++ {
			sum += i
			if i%2 == 0 {
				evenCount++
			}
		}
		fmt.Printf("Сумма чисел от 1 до %d: %d\n", n, sum)
		fmt.Printf("Количество чётных чисел: %d\n", evenCount)

		printed := 0
		fmt.Print("Первые чётные числа: ")
		for i := 1; i <= n; i++ {
			if i%2 != 0 {
				continue
			}
			fmt.Printf("%d ", i)
			printed++
			if printed == 3 {
				break
			}
		}
		fmt.Println()
	}
}
