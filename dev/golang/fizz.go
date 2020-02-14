/* 
- multiple of three, print fizz
- multiple of five, print buzz
- multiple of three and five, print fizz buzz.

*/

package main

import (
	"fmt"
	"os"
	"strconv"
)


func main() {

/* 

Receive 1 argument as input (maximum number)
Starting from 0, perform fizz buzz till maximum number included

*/

	if totalargs := len(os.Args); totalargs != 2 { 
		fmt.Println("Format: <fizz maxnumber>, but only received ", totalargs)
		os.Exit(404)
	}
	
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Exiting")
		os.Exit(404)		
	} else {
		fmt.Println("Max Number is  ",num);
	}

	for i := 0; i <= num; i++ {

		switch i{
		case i % 5 == 0 && i % 3 == 0 :
			fmt.Printf("FizzBuzz %d\n", i)
		case i % 3 == 0 :
			fmt.Printf("Buzz %d\n", i)
		case i % 5 == 0 :
			fmt.Printf("Fizz %d\n", i)
		default:
			fmt.Printf("%d\n",i)
		}
		
		/*
		if  i % 5 == 0 &&  i % 3 == 0 {
			fmt.Println("FizzBuzz -->",i)
		} else if i % 3 == 0 {
			fmt.Println("Fizz -->", i)
		} else if i % 5 == 0 {
			fmt.Println("Buzz -->", i)
		} else {
			fmt.Println(i)
		}
		*/

	}

}
