//Given two integers dividend and divisor, 
// divide two integers without using multiplication, division and mod operator.


/* Combinations

10, 3 : 3 + 3 + 3
10, -3 : -3 - -3 - -3 - -3 - -3
-10, -3
-10, 3

treat both as positive numbers, find the number and switch the sign


*/

package main

import (
	"fmt"
)

func switchInt(orig int) int {

// switch the sign of the number
	return 0-orig


}

func main() {

	var dividend, divisor, quotient, remainder = -7,-3,0,0

	opdividend, opdivisor := dividend, divisor
	if dividend < 0  {
		opdividend = switchInt(dividend)
	}
	if divisor < 0 {
		opdivisor = switchInt(divisor)
	}
	fmt.Printf("Dividend %d, Opdividend %d, Divisor %d, Opdivisor %d\n",dividend, opdividend, divisor, opdivisor)

	for i := 0;  i < opdividend; i += opdivisor {
		if i + opdivisor > opdividend {
			remainder = opdividend - i
			break
		}
		quotient += 1
	}

	if (dividend != opdividend) && (divisor != opdivisor) {
		remainder = switchInt(remainder)
	} else if (dividend == opdividend) && (divisor != opdivisor) {
		quotient = switchInt(quotient)
	} else if (dividend != opdividend) && (divisor == opdivisor) {
		quotient = switchInt(quotient)
		remainder = switchInt(remainder)
	}
	fmt.Println("quotient: ",quotient)
	fmt.Println("remainder: ",remainder)
}


