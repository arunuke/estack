

package main

import (
	"fmt"
	"os" 
	"strconv"
)

func main() {

	var StrLength = len(os.Args)
	fmt.Println(StrLength)
	for i := 0; i < StrLength; i ++ {
		fmt.Println(os.Args[i])

	}
	for i,j:= range os.Args[0:] {
		idx, val := "Index is" + strconv.Itoa(i), "Value is" + j
		// var idx := "Index is" + strconv.Itoa(i)
		//var val := " Value is" + j
		newstr := idx + val
		fmt.Println(newstr)
	}

}

