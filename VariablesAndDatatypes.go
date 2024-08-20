package main

import "fmt"

func variablesAndDatatypes() {
	var intNum int16 = 32767 //Max Integer in 16 bit

	// var intNum int16 = 32767 + 1  -- Throws an Error / you can use int64 but it will take twice the memory.

	var uintNum uint16 = 32767 + 100 // unsigned integer can store double the size of signed integer but only positive numbers.

	var floatNum32 float64 = 1.123456789074 // Output - 1.1234568 Precision is low as compared to float64
	var floatNum64 float64 = 1.123456789074

	fmt.Println(intNum, uintNum, floatNum32, floatNum64)
}