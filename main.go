package main

import (
	"fmt"
)

func main() {
	var input int
	var languages = [4]string{"English", "Spanish", "German", "French"}
	for i := 0; i < len(languages); i++ {
		fmt.Println(i+1, ") ", languages[i])
	}
	fmt.Print("Input Number: ")
	fmt.Scanf("%d", &input)
	fmt.Println()

	var greeting string
	var translations = [4]string{"Hello", "Hola", "Guten Tag", "Bonjour"}

	if input == 0 || input > len(translations) {
		greeting = "Yo"
	} else {
		greeting = translations[input-1]
	}

	fmt.Println(greeting + ", Go!")
}
