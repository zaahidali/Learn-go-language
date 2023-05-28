# Go Practice Document

## GO Language
1. Statically type language
1. Code compiled into binary code.
2. Can translate code for other OS
3. Has great garbage collector for memory management


## PROS:
1. Statically Typed language -> Type checks before compilation --> Variable types are known before compilation like C, C++ & JAVA
2. Efficient memory Management -> so it is fast
3. Has great standard library -> which reduces needs for other external dependencies. It includes packages for HTTP, JSON, file I/O, and more.
4. Cross-Compilation -> It allows code to be compiled for different platforms without much hassle
5. Concurrency -> Go uses goroutines and channels to achieve lightweight and efficient concurrency. Ruby relies on threads and fibers, which are heavier and less efficient compared to goroutines.

## CONS:
1. Smaller community as compared to other languages/frameworks
2. Limited GUI Support
3. Smaller libraries such as Rails has a lot of gems
4. Minimalistic syntax


# Basic Code Syntax
```
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
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

1. `fmt.Println()`
2. `fmt.Print()`
3. `fmt.Printf()`


```
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("Hello, Sparta!")
}
```

```
Output: 
Hello, World! 
Hello, Sparta!
```

```
package main

import "fmt"

func main() {
    fmt.Println("Hello ")
    fmt.Println("Sparta!")
}
```
``` Output: Hello Sparta! ```


```
package main

import "fmt"

func main() {

    fmt.Printf("%s\n", "Hello, Sparta!")
    fmt.Printf("%d\n", 12345)
    fmt.Printf("%g\n", 12.50)
    fmt.Printf("%t\n", true)
}
'\n' means new Line
```
``` 
Output: 
Hello, Sparta!
12345
12.50
true
```


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

## Slice Methods

| Method   | Description                                | Example                                     |
|----------|--------------------------------------------|---------------------------------------------|
| append() | Adds elements to a slice.                  | `slice = append(slice, 4, 5)`              |
| copy()   | Copies elements from one slice to another. | `copy(destination, source)`                |
| DeepEqual() | Compares two slices for equality.       | `isEqual := reflect.DeepEqual(slice1, slice2)` |
| len()    | Returns the length of a slice.             | `length := len(slice)`                      |


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

