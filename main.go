package main

//Задание 3. Вывод ёлочки

import (
	"fmt"
)

func main() {

	var height int
	fmt.Println("Введите высоту ёлочки:")
	fmt.Scan(&height)

	space := " "
	star := "*"
	var line string
	for i := 1; i <= height; i++ {
		line = ""
		for k := height - i; k > 0; k-- {
			line += space
		}
		for k := (i-1)*2 + 1; k > 0; k-- {
			line += star
		}
		fmt.Println(line)
	}

}
