package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure" // Убедись, что эта строка тут есть
)

func main() {
	// В Go мы используем := для краткого объявления и инициализации
	greetings := "Hello world"
	myFigure := figure.NewFigure(greetings, "", true)
	myFigure.Print()
	// Пакет fmt — это стандарт для ввода/вывода
	// fmt.Println(greetings)
	fmt.Println("\nНажмите Enter, чтобы выйти...")
	fmt.Scanln()
}
