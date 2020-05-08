//Given a string containing digits from 2-9 inclusive, return all possible letter 
//combinations that the number could represent.

package main


import (
	"fmt"
)


func main() {

	inStr := "234"
	phoneBoard := make(map[int]string,8)

	phoneBoard[2] = "abc"
	phoneBoard[3] = "def"
	phoneBoard[4] = "ghi"
	phoneBoard[5] = "jkl"
	phoneBoard[6] = "mno"
	phoneBoard[7] = "pqrs"
	phoneBoard[8] = "tuv"
	phoneBoard[9] = "wxyz"
	fmt.Println(phoneBoard)
	fmt.Println(inStr)

	uNumber := len(inStr)
	bStr := make([]byte, uNumber)
	bStr = []byte (inStr)
	fmt.Println(bStr)

	var combiMatrix [27][3] int
}
