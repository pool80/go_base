package main

import "fmt"

func main() {
	//из слайса с днями недели надо скопировать в новый слайс рабочие дни, а из исходного слайса удалить скопированные, чтобы остались только выходные дни.
	//Выведите на экран слайсы с выходными и рабочими днями недели.
	days := []string{"mon", "tue", "wen", "thu", "fri", "sat", "sun"}
	//days := []int{1, 2, 3, 4, 5, 6, 7}
	workDays := make([]string, 5, 7)
	copy(workDays, days)
	days = append(days[5:7], days[7:]...)
	fmt.Println(workDays, days)

	//нужно объединить слайс с выходными днями и слайс с рабочими в один слайс. Выведите на экран итоговый слайс с днями.
	week := append(workDays, days...)
	fmt.Println(week)
}