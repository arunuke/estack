// take int as input, reverse the digits

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	input, _ := strconv.Atoi(os.Args[1])
	fmt.Println("Input",input)
	mult := 10

	arr := make([]int, 10)
	arrlen := 0
	for ; input != 0; arrlen++  {

		arr[arrlen] = input % mult
		input /= 10

	}

	reverse := 0
	mult = 1
	for j := arrlen-1; j >= 0; j-- {

		reverse +=  arr[j] * mult
		mult *= 10
	}
	fmt.Println("Output",reverse)


}
