package main

import "fmt"

func main() {
	var x, y, z, max int
	fmt.Println("Добро пожаловать, мы тут определяем максимальное число из 3 заданных числ")
	fmt.Println("Введи 3 числа: (После каждого числа пробел)")
	fmt.Scan(&x, &y, &z)
	max = x
	if y > max {
		max = y
	}
	if z > max {
		max = z
	}
	fmt.Println("Максимальное число:", max)
}

//Задача 1. Определение максимального числа
//Напишите программу, которая запрашивает у пользователя три целых числа и определяет, какое
//из них является наибольшим. Результат должен выводиться на экран.
//Пример:
//Введите три целых числа: 5 12 8
//Наибольшее число: 12