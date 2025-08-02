package dataTypes

import "fmt"

func BasicDataTypes() {
	var a int = 10
	var b float64 = 3.14
	var c string = "GoLang"
	var d bool = true

	fmt.Println("Basic Data Types:")
	fmt.Println("Integer:", a)
	fmt.Println("Float:", b)
	fmt.Println("String:", c)
	fmt.Println("Boolean:", d)
}

func ArrayOperations() {
	arr := [3]int{1, 2, 3}
	slice := []string{"apple", "banana", "cherry"}

	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)

	 // Map
	 stock := map[string]int{
        "pen":   10,
        "paper": 50,
    }
    fmt.Println("Map:", stock)
    fmt.Println()
}


// Struct definition
type Person struct {
    name string
    age  int
}

func Struct(){
	p := Person{name: "Alice", age: 25}
    fmt.Println("Struct (Person):", p)
    fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
    fmt.Println()

	// Pointer
    var x int = 100
    var ptr *int = &x
    fmt.Println("Pointer:")
    fmt.Println("Address:", ptr)
    fmt.Println("Value at Address:", *ptr)
    fmt.Println()
}

// Interface definition
type Speaker interface {
    Speak()
}

// Struct implementing the interface
type Dog struct{}

func (d Dog) Speak() {
    fmt.Println("Dog says: Woof!")
}

func Interface(){
	var s Speaker = Dog{}
    fmt.Println("Interface:")
    s.Speak()
}

// Function

func AddNumbers(a int, b int) int {
	return a + b
}

// Loop

func PrintNumbers(n int) {
	fmt.Println("Numbers from 1 to", n)
	for i := 1; i <= n; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func CheckEvenOdd(x int) {
	if x%2 == 0 {
		fmt.Printf("%d is Even\n", x)
	} else {
		fmt.Printf("%d is Odd\n", x)
	}
}