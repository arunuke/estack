// Create a slice and then calculate the maximum value among the elements in the slice


package main

import (
	"fmt"
)

func main() {

	slice_of_int := [] int {10, 12,14, 6, 5, 9, 17,25}

	slice_len := len(slice_of_int)
	val := 0
	for i := 0; i < slice_len; i ++ {
		if slice_of_int [i] > val {
			val = slice_of_int [i]
		}

	}
	fmt.Println("Max value is", val);

}
