//Write an interface for converting all upper case to lower case


package main

import (
	"fmt"
	"io"
)

type Capper struct {

	wtr io.Writer

}

func (c *Capper) Write (p []byte) (n int, err error) {

	lenstr := len(p)
	retstr := p
	for i, ch := range p {
		if ch >= 97 && ch <=  122 {
			retstr[i] = ch - 32
		}
	}
	return 
}

func main() {

	teststr := "abcdefghiJKLMNopQR"



}


