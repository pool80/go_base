/*
Напишите функцию contains, которая принимает на вход два параметра: слайс строк a и строку x. 
Функция должна проверять, содержится ли строка x в слайсе a, и возвращать булево значение.

Создайте вариативную функцию getMax, которая находит максимальное целое число из переданных на вход параметров.
Выведите на экран результат вызова функций.
*/

package main

import  "fmt"

func main() {
	contains([]string{"hello", "world"}, "world")
	getMax(5, 6, 7, 35, 5, 66, -4)
}

func contains(a []string, x string) bool{
	var z bool
	for _, word := range a {
		if word == x{
			z = true
		} else {
			z = false
		}
	}
	fmt.Println(z)
	return z
}

func getMax(n...int) int {
	max := n[0]
	for _, i := range n {
		if max <= i {
			max = i
		}
	}
	fmt.Println(max)
	return max
}
