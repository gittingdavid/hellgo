package main

import (
	"fmt"
)

func main() {
	var input int
	var greeting string
	var languages = []string{"English", "Spanish", "German", "French"}
	var translations = []string{"Hello", "Hola", "Guten Tag", "Bonjour"}

	for i := 0; i < len(languages); i++ {
		//fmt.Println(i+1, ") ", languages[i])
		fmt.Print(i+1, ") ")
		fmt.Println(languages[i])
	}

	fmt.Print("Input Number: ")
	fmt.Scanf("%d", &input)
	fmt.Println()

	if input > 0 && input <= len(translations) {
		greeting = translations[input-1]
	} else {
		greeting = "Yo"
	}

	fmt.Println(greeting + ", Go!")
}

func printhelp() {
	fmt.Println("usage: hellgo args1")
	fmt.Println("args: ")
}
