package main

import (
	"fmt"
)

func main() {
	var intArr [3]int32 = [3]int32{1, 2, 3}
	// intArr := [3]int32{1, 2, 3}

	fmt.Println(intArr)
	//intArr[1] = 123
	//fmt.Println(intArr[0])
	//fmt.Println(intArr[1:3])
	//
	//fmt.Println(&intArr[0])
	//fmt.Println(&intArr[1])
	//fmt.Println(&intArr[2])

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("the len is %v with a capacity of %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nthe len is %v with a capacity of %v", len(intSlice), cap(intSlice))

}
