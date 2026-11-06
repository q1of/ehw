package main

import (
	"errors"
	"fmt"
	"math"
)

var ErrInvalidSize = errors.New("неправильный размер")

type Shape interface {
	Area() (float64, error)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() (float64, error) {
	if r.Width <= 0 || r.Height <= 0 {
		return 0, ErrInvalidSize
	}
	return r.Width * r.Height, nil
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() (float64, error) {
	if c.Radius <= 0 {
		return 0, ErrInvalidSize
	}
	return math.Pi * c.Radius * c.Radius, nil
}

func PrintArea(s Shape) error {
	defer fmt.Println("Расчёт завершён")

	area, err := s.Area()
	if err != nil {
		return fmt.Errorf("ошибка: %w", err)
	}

	fmt.Printf("Площадь: %.2f\n", area)
	return nil
}

func showDefer() {
	defer fmt.Println("Второй defer")
	defer fmt.Println("Первый defer")
	fmt.Println("Обычная строка")
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 2},
		Rectangle{Width: 0, Height: 5},
		Circle{Radius: -1},
	}

	for i, s := range shapes {
		fmt.Println("Фигура", i+1)

		err := PrintArea(s)
		if err != nil {
			fmt.Println("Ошибка:", err)
			if errors.Is(err, ErrInvalidSize) {
				fmt.Println("Проверьте размер")
			}
		}

		fmt.Println()
	}

	showDefer()
}
