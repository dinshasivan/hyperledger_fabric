# hyperledger_fabric

# Go-Lang
Go is a programming language created by Google. Go is a fast statically typed, compiled language. It is a general purpose language. Go has built-in testing support. Go is an object-oriented language(in its own way).


## Variable Declaration

```var var_name type```

**Explicit Declaration** 

```var name string```

**Type Inference**

```var name = "value"```

**short variable declaration**

```
    name := "dinsha"
    age := "23"
```

The ```:=``` Operator declares and initialize the varial in one step.

### User Inputs
In Go, you can take user input using the ```fmt.Scan```, ```fmt.Scanln```, or ```fmt.Scanf``` functions.

Using ```fmt.Scanf``` requires a **format specifier** (```%s``` for string, ```%d``` integer etc.)

Reading multi-world or multi-line input you can use ```bufio.Reader```. 
```
    reader := bufio.NewReader(os.Stdin) // Create a buffered reader
```

## Data Types

In Go (Golang), data types are categorized broadly into four groups:

1. Basic types – numbers, strings, booleans

2. Aggregate types – arrays, structs

3. Reference types – slices, maps, pointers, functions, channels

4. Interface types – interfaces

## Format Specifiers 

1. Basic Format Specifiers
    ``` %v``` Defualt format for a value
    ``` %T``` Type of value
    ``` %%``` print a literal %
2. String and Character Format Specifiers
    ``` %s``` String
    ``` %q``` Quoted String
    ``` %c``` Character  
    ``` %U``` Unicode format  
3. Pointer Format Specifier
    ``` %p``` Pointer address

## Array

An array in Go is a fixed-size collection of elements of the same data type. The size of an array is defined at declaration and cannot be changed.

**Declare Array**

```arr := [...]int{10,20,30,40}```

Find array length using **len()** method.

**Slice**

A slice is a dynamic-sized, more flexible alternative to an array.

**Declare Slice and initialize values**

```slice := []string{"Go", "Lang", "Slice"}```

The **make()** function create a slice with a specified length and capacity.

```
    arr := [5]int{1, 2, 3, 4, 5}

    slice := arr[1:4] 
```
The **append()** function in Go is used to add elements to a slice. It dynamically increases the slice's capacity if needed.

## Loops
Go has only one looping construct: the for loop. However, it can be used in different ways to achieve while and do-while loop behavior.

**Basic ```for``` loop**
```
    for i := 0; i < 5; i++{
        //statements
    }
```

**```for``` loop as ```while``` loop**
```
    x := 0
	for x < 5{
	    /statement
	    // increment x 
	}
```

## Conditionals & Booleans
Go provides standard conditional statements like ```if```, ```else if```, ```else```, and ```switch```. Boolean values (```true``` / ```false```) are essential in controlling program flow.

## Functions

Functions in Go help in modularizing code and improving reusability. They are defined using the func keyword.

* Use parameteers for input and return values for output
* Go supports multiple retun values.
* Functions can be assigned to varables and passed as arguments


### Pointers
A pointer in Go is a variable that stores the memory address of another variable. Instead of holding an actual value, it "points" to where the value is stored in memory.

Use pointers (```*T```) to modify the original variable.

### Struct
A struct is a collection of fields. It groups related values under a single type.

## Interface

An interface in Go is a type that specifies a set of method signatures. A type implements an interface if it defines the methods declared by the interface. Go interfaces provide polymorphism without explicit inheritance.

