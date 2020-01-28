/* 

Find even ended numbers within a given range

Get a range of numbers as input argument.

Calculate product of all numbers

Find even ended numbers within that list 

Enhance : Instead of a number range, accept number of digits

*/


package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {

	// Incoming arguments

	if arglen := len(os.Args); arglen != 3 {
		fmt.Println("Need 2 arguments. <progname> <lower range> <higher range>", arglen)
		os.Exit(404) 
	}

	// Identify upper and lower bounds
	// Enhance: function takes all input strings and returns as many values as determined
	lower, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("strconf error", err)
	}
	upper, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("strconf error", err)
	}



	// Run loop

	for i := lower; i <= upper; i ++ {
		for j := i; j <= upper; j ++ {
			// Multiply numbers
			product := i * j

			// Convert product to string
			strproduct := fmt.Sprintf("%d",product);

			// Determine if the product is even-ended
			// Check if [0] and [n] are the same character
			strlen := len(strproduct)

			// if even-ended, print the number
			if strproduct[0] == strproduct [strlen-1] {
				fmt.Printf("%d is even ended; From %d and %d\n",product, i, j)
			}


		}
	}


}
