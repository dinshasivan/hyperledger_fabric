package main

import "fmt"

func main()  {
	fmt.Println("Indroduction to Go")
	fmt.Println("Hello world!")

	var name string
	fmt.Println("Enter your name:")
	fmt.Scan(&name)
	fmt.Println("Hello %s good morning \n", name)
	
}