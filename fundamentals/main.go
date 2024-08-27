package main

import "fmt"

func main() {
	var intNum int16 = 32767
	// hits int limit and goes negative
	//intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var results float32 = floatNum32 + float32(intNum32)
	fmt.Println(results)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var myString string = "Hello World"
	fmt.Println(myString)

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = true
	var myBoolean2 bool = false

	fmt.Println(myBoolean)
	fmt.Println(myBoolean2)

	myVar := "test"
	fmt.Println(myVar)

	var var1, var2 int = 1, 2
	fmt.Println(var1 + var2)

	var3, var4 := 3, 4

	fmt.Println(var3 + var4)

	const myConst string = "const example"

	fmt.Println(myConst)
}