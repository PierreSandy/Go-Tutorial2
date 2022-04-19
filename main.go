package main

import (
	"fmt"
	"strconv"
	"reflect"
	"log"
	
)


//  Converting  a string into an integer 

func main () {


	// Atoi()Function - to convert the string into an integer value. Atoi stands for ASC to integer. the Atoi() function returns two values: the result of the converson and
	// any error 
	// syndax func Atoi (s string) (int, error)

	strVar := "100"
	intVar, err := strconv.Atoi(strVar)  // converts string 
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
	log.Println(intVar, err, reflect.TypeOf(intVar))

	//      ARITHMETIC OPERATIONS
	// + 	Addition 	x + y 	Sum of x and y
	// - 	Subtraction 	x - y 	Subtracts one value from another
	// * 	Multiplication 	x * y 	Multiplies two values
	// / 	Division 	x / y 	Quotient of x and y
	// % 	Modulus 	x % y 	Remainder of x divided by y
	// ++ 	Increment 	x++ 	Increases the value of a variable by 1
	// -- 	Decrement 	x-- 	Decreases the value of a variable by 1

	var x, y = 1, 2

	fmt.Printf(" x + y = %d\n",x+y) // 1 + 2
	fmt.Printf(" x - y = %d\n", y-x) // 2 - 1
	fmt.Printf(" x * y = %d\n", x*y) //1 * 2
	fmt.Printf("x / y = %d\n", x/y)// 1 / 2
	fmt.Printf("x mod y = %d\n", x%y) // Modulus operation 

	x++
	fmt.Printf("x++ = %d\n", x) // increments number by 1

	y--
	fmt.Printf("x-- = %d\n", y) // decreases number by 1


    //ASSIGNMENT OPERATORS 
	// x = y 	Assign 	x = y
	// x += y 	Add and assign 	x = x + y
	// x -= y 	Subtract and assign 	x = x - y
	// x *= y 	Multiply and assign 	x = x * y
	// x /= y 	Divide and assign quotient 	x = x / y
	// x %= y 	Divide and assign modulus 	x = x % y

	x = y
	fmt.Println(" =", y) // equal to 1

	x += y
	fmt.Println("+=", x) // equal to 2

	x -= y
	fmt.Println("-=", x)

	x *= y
	fmt.Println("*=", x)

	x = 150
	x /= y
	fmt.Println("/=", x)

	x = 40
	x %= y
	fmt.Println("%=", x)

    //COMPARISON OPERATORS
	// == 	Equal 	x == y 	True if x is equal to y
	// != 	Not equal 	x != y 	True if x is not equal to y
	// < 	Less than 	x < y 	True if x is less than y
	// <= 	Less than or equal to 	x <= y 	True if x is less than or equal to y
	// > 	Greater than 	x > y 	True if x is greater than y
	// >= 	Greater than or equal to 	x >= y 	True if x is greater than or equal to y
    

	//comparing x = 1 & y = 2

    
	fmt.Println(x == y)



	}