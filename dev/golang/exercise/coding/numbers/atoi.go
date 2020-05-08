// implement atoi
// discard any leading white spaces
// initial + or - followed by as many numerical digits
// any additional characters beyond the numbers are ignored


package main

import (
	"fmt"
	"os"
)


func main() {

	lstr := os.Args[1]
	i,j := 0,0

	slen := len(lstr)
	for i = 0; i < slen; i++ {
		if lstr[i] >= 48 &&  lstr[i] <=57 {
			fmt.Println("number found at: ",i) 
			for j = i; j < slen; j ++ {
				if lstr[j] < 48 ||  lstr[j] > 57 {
					fmt.Println("number ends at: ", j) 
					break
				} 
			}
			fmt.Println("i and j : ", i, j)
			break
		}
	}
	fmt.Println("string is", lstr[i:j])
}
