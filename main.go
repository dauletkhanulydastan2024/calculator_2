package main

import "fmt"

func main() {
	var a, b int
	var operator string
	fmt.Println("a: ")
	fmt.Scan(&a)
	fmt.Println("b: ")
	fmt.Scan(&b)
	fmt.Println("operator: ")
	fmt.Scan(&operator)
	fmt.Println("operator:", operator)
	// Используя логические конструкции `if/else` или `switch`
	// имплементируйте логику вызова соответствующей функции
	// в зависимости от значения переменной `operator`.
	// Если введенный пользователем оператор неизвестен,
	// то вывести сообщение "Неизвестный оператор".
	var result int
	// ПИШЕМ КОД ТУТ
	if operator == "+" {
		result = plus(a, b)
	} else if operator == "-" {
		result = minus(a, b)
	} else if operator == "*" {
		result = multyply(a, b)
	} else if operator == "/" {
		result = divide(a, b)
	} else {
		fmt.Println("???")
	}
	fmt.Printf("%d %s %d = %d", a, operator, b, result)
}

// Ниже этой строчки имплементируйте функцию `Add`,
func plus(a, b int) int {
	return a + b
}

// которая выполняет произведение двух чисел и возвращает результат
// Ниже этой строчки имплементируйте функцию `Subtract`,
func minus(a, b int) int {
	return a - b
}

// которая выполняет вычитание двух чисел и возвращает результат
// Ниже этой строчки имплементируйте функцию `Multiply`,
func multyply(a, b int) int {
	return a * b
}

// которая выполняет произведение двух чисел и возвращает результат
// Ниже этой строчки имплементируйте функцию `Divide`,
func divide(a, b int) int {
	return a / b
}

// которая выполняет деление двух чисел и возвращает результат
