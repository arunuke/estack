//From the golang course

//Calcuate the mean value of two numbers


package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var sum int 
	numbers := len(os.Args)
	if  numbers  < 3 {
		fmt.Println("# of arguments is less than 3; Format: mean <arg1> <arg2>;  Exiting")	
		os.Exit(404)

	} else {
		fmt.Println("# of arguments is",len(os.Args)) 
	}

	for i := 1; i < numbers; i ++ {
	//	fmt.Println(os.Args[i])	
		if tmpsum, err := strconv.Atoi(os.Args[i]); err == nil {
			sum = sum + tmpsum
		} else {
			fmt.Println("Error in Atoi",err)
			os.Exit(404)
		}	
	}
	avg := sum / (numbers - 1)

	fmt.Println("Average is", avg)

}
