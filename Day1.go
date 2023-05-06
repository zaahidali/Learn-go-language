package main

import "fmt"

func main() {
	fmt.Println("############### DAY 1############################")
	var number1, number2 float64

	// 1 ---- prompt 2 numbers and calculate their Division
	fmt.Println("Enter 2 numbers for division: ")
	_, err := fmt.Scan(&number1, &number2)
	if checkError(err) {
		fmt.Println("Invalid input type data entered")
		return
	}
	divideNumbers(number1, number2)

	// 2 final
	fmt.Println("\nEnter 2 numbers to swap their values")
	_, err1 := fmt.Scan(&number1, &number2)
	if checkError(err1) {
		fmt.Println("Only integer & float numbers values swaps are allowed")
		return
	}
	swapValues(number1, number2)

	// 3 promt user to ask for name, age, dob and height in ft
	simpleConverter()
}

func divideNumbers(number1 float64, number2 float64) {
	if number2 == 0 {
		fmt.Println("Denominator cannot be zero")
	} else {
		divResult := number1 / number2
		fmt.Printf("The Result of %v/%v is: %v\n", number1, number2, divResult)
	}
}

func swapValues(number1 float64, number2 float64) {
	fmt.Println("\nBefore swapping")
	fmt.Printf("number1 = %v number2 = %v", number1, number2)
	number2, number1 = number1, number2
	fmt.Printf("\nAfter Swapping \nnumber1 = %v number2 = %v\n", number1, number2)
}

func simpleConverter() {
	var name string
	var age int
	var height float32

	fmt.Print("\nEnter your name: ")
	_, err := fmt.Scanf("%s\n", &name)
	if checkError(err) {
		fmt.Println("Invalid input for name")
		return
	}

	fmt.Print("Enter your age in years: ")
	_, err = fmt.Scanf("%d\n", &age)
	if checkError(err) {
		fmt.Println("Invalid input for age")
		return
	}

	fmt.Print("Enter your height in feet: ")
	_, err = fmt.Scanf("%g\n", &height)
	if checkError(err) {
		fmt.Println("Invalid input for height")
		return
	}

	if age <= 0 || age > 160 {
		fmt.Println("Invalid age")
		return
	}

	if height <= 0 || height >= 15 {
		fmt.Println("Invalid height")
		return
	}

	fmt.Printf("\nName: %s\nAge in years: %d\nAge in days: %d\nHeight in feet: %g\nHeight in meters: %g\n",
		name, age, age*365, height, height*0.3048)
}

func checkError(err error) bool {
	return err != nil
}