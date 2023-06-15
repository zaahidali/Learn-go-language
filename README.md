# Go Programming Essentials

## Table of Contents

- [Overview](#overview)
- [Introduction to Go](#introduction-to-go)
- [Pros and Cons](#pros-and-cons)
- [Basic Code Syntax](#basic-code-syntax)
- [Go Print Statements](#go-print-statements)
- [Go Format Specifiers](#go-format-specifiers)
- [String in Go](#string-in-go)
- [String methods](#string-methods)
- [Slice in Go](#slice-in-go)
- [Slice methods](#slice-methods)
- [Array vs Slice](#go-array-vs-slice)
- [Go map](#go-map)

## Overview

Welcome to the ***Go Practice Document***, where we'll explore the features and benefits of the Go programming language. This document aims to provide a professional overview of Go and highlight its important aspects.


## Introduction to Go

Go is a statically typed programming language that compiles code into binary executables. It offers several key advantages that make it a popular choice among developers:

## **Pros and Cons**

- **Pros**:
  - **Statically Typed Language**:
    - Type checks before compilation
    - Variable types are known before compilation like C, C++ & Java
  - **Efficient Memory Management**:
    - Fast and reliable
    - Great garbage collector for memory management
  - **Extensive Standard Library**:
    - Includes packages for HTTP, JSON, file I/O, and more
    - Reduces the need for external dependencies
  - **Cross-Compilation**:
    - Code can be compiled for different platforms without hassle
  - **Concurrency**:
    - Uses goroutines and channels for lightweight and efficient concurrency

- **Cons**:
  - **Smaller Community** compared to other languages/frameworks
  - **Limited GUI Support**
  - **Smaller Libraries** compared to languages like Ruby
  - **Minimalistic Syntax**


## Basic Code Syntax

To provide a taste of Go, here's an example of the classic ```"Hello, World!"``` program in Go:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Output:
```
Hello, World!
```

## Data Types

| Type     | Description                                | Examples          |
|----------|--------------------------------------------|-------------------|
| int      | Integer numbers.                           | 7123, 0, -5, 7023 |
| float    | Numbers with decimal points.               | 20.2, 500.123456, -34.23 |
| complex  | Complex numbers.                           | 2+4i, -9.5+18.3i  |
| string   | Sequence of characters.                    | "Hello World!", "1 is less than 2" |
| bool     | Either true or false.                      | true, false       |
| byte     | A byte (8 bits) of non-negative integers. | 2, 115, 97        |
| rune     | Used for characters. Internally used as 32-bit integers. | 'a', '7', '<' |


## Go Print Statements

The Go programming language provides several print statement options that allow you to output text and values to the console. Here are three commonly used print statements in Go:

1. `fmt.Println()`: This statement prints the given arguments to the console, adding a new line after each print.

2. `fmt.Print()`: This statement prints the given arguments to the console without adding a new line. Subsequent prints will be displayed on the same line.

3. `fmt.Printf()`: This statement provides more control over the output by using format specifiers. It allows you to format and print values based on their type.

### Example 1: `fmt.Println()`
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("Hello, Sparta!")
}
```
Output:
```
Hello, World!
Hello, Sparta!
```

### Example 2: `fmt.Print()`
```go
package main

import "fmt"

func main() {
    fmt.Print("Hello ")
    fmt.Print("Sparta!")
}
```
Output:
```
Hello Sparta!
```

### Example 3: `fmt.Printf()`
```go
package main

import "fmt"

func main() {
    fmt.Printf("%s\n", "Hello, Sparta!")
    fmt.Printf("%d\n", 12345)
    fmt.Printf("%g\n", 12.50)
    fmt.Printf("%t\n", true)
}
```
Output:
```
Hello, Sparta!
12345
12.50
true
```

In the last example, we used format specifiers such as `%s` for strings, `%d` for integers, `%g` for floating-point numbers, and `%t` for booleans. The `\n` represents a new line character, which adds a line break after each print statement.

These print statements are essential for displaying output and debugging in Go, allowing you to communicate with the user and inspect the values of variables during program execution.


# Go Format Specifiers

| Type     | Format Specifier |
|----------|-----------------|
| Integer  | %d              |
| Float    | %g              |
| String   | %s              |
| Boolean  | %t              |
| Character | %c             |
| Hexadecimal | %x           |
| Scientific notation | %e  |
| Octal | %o                 |
| Pointer | %p              |
| Width | %*d               |



# String in Go

- A string in Go is a **sequence of bytes**.
- Strings are **enclosed** in double quotes (`"like this"`) or backticks (`` `like this` ``).
- Strings in Go are **immutable**, meaning they cannot be changed once created.
- Go provides a **rich set of string manipulation functions** in the `strings` package for tasks like **searching, replacing, splitting**, and more.
- Strings can be **compared** using the `==` or `!=` operators for **equality or inequality**.

## String Methods

| Method | Description | Example |
| ------ | ----------- | ------- |
| `Compare()` | Compares two strings lexicographically | `"abc".Compare("abd") // returns -1"` |
| `Contains()` | Returns `true` if a string contains another string | `"abcdef".Contains("cd") // returns true"` |
| `ToUpper()` | Converts a string to uppercase | `"hello world".ToUpper() // returns "HELLO WORLD"` |
| `ToLower()` | Converts a string to lowercase | `"Hello World".ToLower() // returns "hello world"` |
| `Replace()` | Replaces all occurrences of a substring with another substring | `"hello world".Replace("world", "everyone") // returns "hello everyone"` |
| `Split()` | Splits a string into an array of substrings based on a separator string | `"hello,world".Split(",") // returns ["hello", "world"]"` |
| `Join()` | Joins an array of substrings into a single string using a separator string | `strings.Join([]string{"hello", "world"}, ",") // returns "hello,world"` |

## String Formatting

| Specifier | Description | Example |
| --------- | ----------- | ------- |
| `%s` | Formats a string | `fmt.Sprintf("Name: %s", name)` |
| `%d` | Formats an integer | `fmt.Sprintf("Age: %d", age)` |
| `%f` | Formats a floating-point number | `fmt.Sprintf("Price: %.2f", price)` |
| `%t` | Formats a boolean value | `fmt.Sprintf("Is available: %t", isAvailable)` |

## Example Code

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    // Define a string variable
    message := "hello world"

    // Compare two strings
    result := strings.Compare("abc", "abd")
    fmt.Println(result) // Output: -1

    // Check if a string contains another string
    result = strings.Contains(message, "world")
    fmt.Println(result) // Output: true

    // Convert a string to uppercase
    result = strings.ToUpper(message)
    fmt.Println(result) // Output: HELLO WORLD

    // Convert a string to lowercase
    result = strings.ToLower(message)
    fmt.Println(result) // Output: hello world

    // Replace all occurrences of a substring with another substring
    result = strings.Replace(message, "world", "everyone", -1)
    fmt.Println(result) // Output: hello everyone

    // Split a string into an array of substrings
    resultArray := strings.Split(message, " ")
    fmt.Println(resultArray) // Output: [hello world]

    // Join an array of substrings into a single string
    result = strings.Join(resultArray, ",")
    fmt.Println(result) // Output: hello,world
}
```

## Slice in Go

- A **slice** in Go is a **variable-size sequence** of elements that provides a more **flexible and powerful alternative** to arrays.
- Slices are built **on top of arrays** and provide a **dynamic view** into the underlying array.
- Unlike arrays, slices can **grow or shrink** in size as needed, allowing for more **efficient memory usage**.
- Slices are represented by a **three-part structure**: a **pointer to the underlying array**, a **length** indicating the number of elements in the slice, and a **capacity** representing the maximum number of elements the slice can hold.
- Slices are widely used in Go to work with **collections of data**, such as lists, buffers, or dynamic arrays.

### Example: Working with Slices

```go
package main

import "fmt"

func main() {
    // Create a slice
    numbers := []int{1, 2, 3, 4, 5}

    // Accessing elements of a slice
    fmt.Println(numbers[0]) // Output: 1
    fmt.Println(numbers[2]) // Output: 3

    // Modifying elements of a slice
    numbers[1] = 10
    fmt.Println(numbers) // Output: [1 10 3 4 5]

    // Slicing a slice
    slice := numbers[2:4]
    fmt.Println(slice) // Output: [3 4]
}
```

## Slice Methods

| Method   | Description                                | Example                                     |
|----------|--------------------------------------------|---------------------------------------------|
| append() | Adds elements to a slice.                  | `slice = append(slice, 4, 5)`              |
| copy()   | Copies elements from one slice to another. | `copy(destination, source)`                |
| DeepEqual() | Compares two slices for equality.       | `isEqual := reflect.DeepEqual(slice1, slice2)` |
| len()    | Returns the length of a slice.             | `length := len(slice)`                      |

# Go Array vs Slice

| Feature               | Array                            | Slice                              |
|-----------------------|----------------------------------|------------------------------------|
| Definition            | Fixed-size sequence of elements  | Variable-size sequence of elements |
| Syntax                | `var arr [size]Type`             | `var slc []Type`                   |
| Length                | Fixed, cannot change             | Can change, dynamic                |
| Capacity              | Same as length                   | Can be larger than the length      |
| Indexing              | `arr[i]`                         | `slc[i]`                           |
| Creation              | `arr := [3]int{1, 2, 3}`         | `slc := []int{1, 2, 3}`            |
| Resizing              | Not possible                     | Possible using `append` or `copy`  |
| Creating from an array| N/A                              | `arr[start:end]`                   |
| Underlying structure  | Plain array                      | Reference to an array              |
| Memory usage          | Less (in-place storage)          | More (reference + metadata)        |



- **When passing an array to a function**, Go makes a copy of the array, and the function works on the copy. So, any modifications to the array inside the function do not affect the original array in the calling function.
- **On the other hand, when passing a slice to a function**, Go passes a reference to the underlying array, which means that the function can modify the original array.

Here's an example demonstrating the differences:

```go
package main

import (
	"fmt"
)

func modifyArray(arr [3]int) {
	arr[0] = 100
	fmt.Println("Inside modifyArray:", arr)
}

func modifySlice(slc []int) {
	slc[0] = 100
	fmt.Println("Inside modifySlice:", slc)
}

func main() {
	arr := [3]int{1, 2, 3}
	slc := []int{1, 2, 3}

	fmt.Println("Before modifyArray:", arr)
	modifyArray(arr)
	fmt.Println("After modifyArray:", arr)

	fmt.Println("\nBefore modifySlice:", slc)
	modifySlice(slc)
	fmt.Println("After modifySlice:", slc)
}
```
```
OUTPUT

Before modifyArray: [1 2 3]
Inside modifyArray: [100 2 3]
After modifyArray: [1 2 3]

Before modifySlice: [1 2 3]
Inside modifySlice: [100 2 3]
After modifySlice: [100 2 3]

```

In summary, when passing an array to a function, you're working with a copy, while when passing a slice, you're working with a reference to the original data. This is the primary difference when passing arrays and slices to functions in Go.

# Go Map

**Go maps** are a built-in data structure that associate keys with values. They are similar to dictionaries in Python or hash tables in other programming languages. Some key features of Go maps include:

- **Dynamic size**: Maps in Go can grow or shrink in size as needed.
- **Unordered**: The order in which elements are stored in a map is not guaranteed, and iterating over a map may yield different orders.
- **Reference type**: Maps are reference types, which means that when you assign a map to another variable, they both point to the same underlying data. Modifying one will affect the other.
- **Zero value**: The zero value of a map is `nil`. A `nil` map has no keys, nor can keys be added.

Here's a summary of common map operations in Go:

- **Create a map**: To create a map, you can use the `make` function or initialize it with key-value pairs.
- **Add/Update an entry**: To add or update a key-value pair in the map, use the assignment operator (`=`).
- **Delete an entry**: To remove a key-value pair from the map, use the `delete` function.
- **Get an entry**: To retrieve the value of a key in the map, use the index operator (`[]`).
- **Check if key exists**: To verify if a key is in the map, use the `, ok` idiom when getting an entry.
- **Length of the map**: To get the number of entries in the map, use the `len` function.
- **Iterate over the map**: To loop through the map's key-value pairs, use the `range` keyword in a `for` loop.


| Operation                | Description                            | Example                                         |
|--------------------------|----------------------------------------|-------------------------------------------------|
| Create a map             | Initialize a new map                   | `ages := make(map[string]int)`                  |
| Create a map with values | Initialize a map with key-value pairs  | `ages := map[string]int{"John": 30, "Jane": 28}`|
| Add an entry             | Add a key-value pair to the map        | `ages["Alice"] = 25`                            |
| Update an entry          | Update the value of a key in the map   | `ages["John"] = 31`                             |
| Delete an entry          | Remove a key-value pair from the map   | `delete(ages, "John")`                          |
| Get an entry             | Retrieve the value of a key in the map | `age := ages["John"]`                           |
| Check if key exists      | Verify if a key is in the map          | `age, ok := ages["John"]; if ok { ... }`        |
| Length of the map        | Get the number of entries in the map   | `count := len(ages)`                            |
| Iterate over the map     | Loop through the map's key-value pairs | `for key, value := range ages { ... }`          |

