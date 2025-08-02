package main

import (
	"fmt"
	"os"

	"Go_Lang_Basic/dataTypes"
)

func main() {
	fmt.Println("Introduction to Go")
	fmt.Println("Hello world!")

	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Hello %s, good morning!\n", name)

	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument to execute a specific function.")
		return
	}

	switch os.Args[1] {
	case "1":
		dataTypes.BasicDataTypes()
	case "2":
		dataTypes.ArrayOperations()
	case "3":
		dataTypes.Struct()
	case "4":
		dataTypes.Interface()
	case "5":
		result :=dataTypes.AddNumbers(5,7)
		fmt.Println(result);
	case "6":
		dataTypes.PrintNumbers(10)
	case "7":
		dataTypes.CheckEvenOdd(8)

	default:
		fmt.Println("Invalid option. Use 1 or 2.")
	}
}
