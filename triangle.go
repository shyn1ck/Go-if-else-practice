package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Println("Добро пожаловать, мы тут определяем правильност треугольника")
	fmt.Println("Введи длину x : ")
	fmt.Scan(&x)
	fmt.Println("Введи длину y : ")
	fmt.Scan(&y)
	fmt.Println("Введи длину z : ")
	fmt.Scan(&z)
	if x+y <= z || y+z <= x || x+z <= y {
		fmt.Println("Такого треугольника не существует!")
	} else {
		fmt.Println("Такой треугольник имеет место быть")
	}
}

//2. Проверка треугольника
//Напишите программу, которая запрашивает у пользователя длины трех сторон треугольника
//и проверяет, может ли такой треугольник существовать. Треугольник существует, если
//сумма длин любых двух его сторон больше длины третьей стороны. Результат проверки
//выводится на экран.