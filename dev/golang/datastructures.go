//Validating some arrays, maps, slices and operations on them

package main

import (

	"fmt"
)

func main() {

//Declare an array

	var  IntArr [5]int

	for i := 0; i <  len(IntArr); i ++ {
		fmt.Printf("Index is %d and value is %d\n ", i, IntArr[i]) 
	}
	fmt.Println("Value of array is", IntArr)

	var StrSlice  []string

}
