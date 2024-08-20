package main

import "fmt"

func variablesAndDatatypes() {
	var intNum int16 = 32767 //Max Integer in 16 bit

	// var intNum int16 = 32767 + 1  -- Throws an Error / you can use int64 but it will take twice the memory.

	var uintNum uint16 = 32767 + 100 // unsigned integer can store double the size of signed integer but only positive numbers.

	var floatNum32 float64 = 1.123456789074 // Output - 1.1234568 Precision is low as compared to float64
	var floatNum64 float64 = 1.123456789074

	fmt.Println(intNum, uintNum, floatNum32, floatNum64)


	//Dividing an integer and a float will cause an error as both datatypes are different.
	// fmt.Println(intNum/floatNum32)

	//Type Conversion is Required
	fmt.Println(uintNum + uint16(floatNum32))

	var string string = "Hi"
	fmt.Println(string)

	//You do not need to initialize the variable everythime
	//In this case it will set it automatically to 0 and for string it will be empty string ''
	var integer int
	fmt.Println(integer)

	//Data type will be automatically inferred
	var integer2 = 2
	integer3 := 3
	fmt.Println(integer2, integer3)

	//use const in something which you dont want value to change aferwards.
	const pi = 3.1415
	fmt.Println(pi)
}