package main

import "fmt"

func main() {
	// В Go мы используем := для краткого объявления и инициализации
	greetings := "Hello world"
	// Пакет fmt — это стандарт для ввода/вывода
	fmt.Println(greetings)
	fmt.Println("\nНажмите Enter, чтобы выйти...")
	fmt.Scanln()
}
