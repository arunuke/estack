/*

Leetcode : Longest substring with no repeat characters

*/


package main

import (
	"fmt"
	"os"
)

func LongestSubstr(s string) {

	// Find the longest substring without repeated character
	// Assume string only has 
	slen := len(s) // length of string stored in slen
	flagmap := make([]bool, 26) //ascii 97 through 22. a to z
	nilmap := make([]bool, 26) 
	i, j, curlen, newlen, start, end := 0,0,0,0,0,0


	for ; i < slen; i ++ {
		for j = i ; j < slen; j ++ {
			idx := 122-s[j]
			if flagmap[idx] == false { // first appearance of character
				flagmap[idx] = true // set map to positive and keep going
			} else { // we have hit the same character
				flagmap = nilmap // reset all values in the flagmap
				newlen = j-i
				if curlen < newlen { // new lengthier substring detected
					curlen = newlen
					start = i
					end = j
				}
				break
			}
		}
	}
	fmt.Printf("\n curlen %d, start %d, end %d\n",curlen, start, end)

}

func IsAnagram(s1 string, s2 string) bool {

	rev1 := SortString([]byte(s1))
	rev2 := SortString([]byte(s2))

	if rev1 == rev2 {
		return true
	} else {
		return false
	}

}

func SortString(s1 []byte) string {

	slen := len(s1) // length of string
	for i := 0; i < slen; i ++ {
		for j :=0;  j < slen; j++  { 
			if s1[i] > s1 [j] {
				tmp := s1[i]
				s1[i] = s1[j]
				s1[j] = tmp
			}
		}
	}
	return (string)(s1)

}

func main() {

	// args[1] is the string

	if len(os.Args) != 3 {
		fmt.Println("Needs two strings as argument. <executable> <string1> <string2>")
	}

	LongestSubstr(os.Args[1])
	if IsAnagram(os.Args[1], os.Args[2]) == true {
		fmt.Println("Anagrams")
	} else {
		fmt.Println("Not Anagrams")
	}
}

