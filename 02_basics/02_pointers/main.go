package main

import (
	"fmt"
	"math"
)

func main() {
	//объявите переменную A как указатель на int;
	//переменную B — целочисленную с произвольным значением;
	//присвойте в переменную A указатель на целочисленную переменную B и выведите на экран значение путем разыменовывания указателя A;
	var a *int
	b := 4
	a = &b
	fmt.Println(*a)

	//присвойте целочисленной переменной B новое произвольное значение через указатель A и выведите на экран новое значение переменной B.
	x := 8
	a = &x
	b = *a
	fmt.Println(b)

	//найдите радиус окружности, если её длина равна 35.
	var p float64 = 35
	pi :=  math.Pi
	rad := p / (2 * pi)
	//радиус окружности R объявите как указатель на float64.
	var r *float64 = &rad
	//вычислите площадь круга, используя при расчёте разыменовывание указателя R.
	s := pi * math.Pow(*r, 2)
	fmt.Println(math.Round(s * 100) / 100)
}