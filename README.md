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
