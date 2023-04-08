package main

import "fmt"

func main() {
	//Arithmetic operator
	fmt.Println("Arithmetic Operator Example")
	p := 5
	q := 4
	//addition
	result := p + q
	fmt.Println("Sum of p+q : ", result)
	//subtraction
	result1 := p - q
	fmt.Println("Sub of p-q : ", result1)
	//multiplication
	result2 := p * q
	fmt.Println("Multiplication of p*q : ", result2)
	//division
	result3 := p / q
	fmt.Println("Division p/q : ", result3)
	//Modulus
	result4 := p % q
	fmt.Println("Modulus p % q : ", result4)

	fmt.Println()
	fmt.Println("Relational Operator Example")
	a := 34
	b := 20
	//equal to operator ==
	result5 := a == b
	fmt.Println("is a == b : ", result5)
	//not equal to operator !=
	result6 := a != b
	fmt.Println("is a != b : ", result6)
	// less than Operator <
	result7 := a < b
	fmt.Println("is a < b : ", result7)
	//Greater than Operator >
	result8 := a > b
	fmt.Println("is a > b : ", result8)
	//Less than equal to Operator <=
	result9 := a <= b
	fmt.Println("is a <= b : ", result9)
	//Greater than equal to Operator >=
	result10 := a >= b
	fmt.Println("is a >= b : ", result10)

	fmt.Println()
	fmt.Println("Logical Operator Example")
	c := 23
	d := 34

	//Logical And && example -it has to ture all condition
	if c != d && c <= d {
		fmt.Println("True")
	}
	// Logical or || example - it has to ture at least one condition
	if c == d || c != d {
		fmt.Println("True")
	}
	// Logical Not ! Example - it has to ture accordingly condition
	if !(c == d) {
		fmt.Println("True")
	}

	fmt.Println()
	fmt.Println("Bitwise Operator Example ")
	e := 34
	f := 20
	//Bitwise And &
	/* Table
	   a  |  b  | and
	   0  |  0  | 0
	   0  |  1  | 0
	   1  |  0  | 0
	   1  |  1  | 1

	*/
	result11 := e & f
	fmt.Printf("Result of e & f = %d\n", result11)
	// Bitwise or
	/* Table
	   a  |  b  | and
	   0  |  0  | 0
	   0  |  1  | 1
	   1  |  0  | 1
	   1  |  1  | 1

	*/
	result12 := e | f
	fmt.Printf("Result of e | f = %d\n", result12)
	// Xor operator ^
	/* Table
	   a  |  b  | and
	   0  |  0  | 0
	   0  |  1  | 1
	   1  |  0  | 1
	   1  |  1  | 0

	*/
	result13 := e ^ f
	fmt.Printf("Result of e ^ f = %d\n ", result13)
	// Left shift <<
	result14 := e << 1
	fmt.Printf("Result of e << 1 = %d\n ", result14)
	// Right shit >>
	result15 := e >> 1
	fmt.Printf("Result of e >> 1 = %d\n ", result15)
	//And NOT &^
	result16 := e &^ f
	fmt.Printf("Result of e &^ f = %d ", result16)
}
