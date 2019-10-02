package main

import (
	"fmt"
)

func main() {
	var input int
	var greeting string
	var languages = [4]string{"English", "Spanish", "German", "French"}
	var translations = [4]string{"Hello", "Hola", "Guten Tag", "Bonjour"}

	for i := 0; i < len(languages); i++ {
		//fmt.Println(i+1, ") ", languages[i])
		fmt.Print(i+1, ") ")
		fmt.Println(languages[i])
	}
	fmt.Print("Input Number: ")
	fmt.Scanf("%d", &input)
	fmt.Println()

	if input == 0 || input > len(translations) {
		greeting = "Yo"
	} else {
		greeting = translations[input-1]
	}

	fmt.Println(greeting + ", Go!")
}
