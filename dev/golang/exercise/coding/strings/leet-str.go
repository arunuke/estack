/*

1. Implement strStr

Return the index of the first occurrence of needle in haystack, 
or -1 if needle is not part of haystack.

2. Find longest palindromic substring


*/

package main

import (
	"fmt"
	"os"
)

func findSubstr(s1 string, s2 string) int {

	slen1, slen2 := len(s1), len(s2)
	i, j := 0,0
	fmt.Println("s1 and s2 lengths:", slen1, slen2)

	for i = 0; i < slen1; i ++ {
		if s1[i] == s2[0] {
			fmt.Println("Found a match for first character  at",i) 
			for j = 1; j < slen2; j++ {
				//move character by character
				fmt.Println("Entered secondary loop")
				if s1[i+j] !=  s2[j] {
					fmt.Println("mismatch in secondary loop")
				// mismatched character - break out of this loop	
				break
				}
			}
			if j == slen2 {
			// breaking out of second loop and j is full length
			// found substring
				fmt.Println("returning index  ",i)
				return i
			}
		}
	}
	return -1
}

func isPalindrome(s1 string) bool {
// if s1 is a palindrome, return true.  else, return false 

// QN: how to find if palindrome without using string length?

	slen :=  len(s1)
	for i := 0; i <= slen/2; i ++ {
		if s1[i] != s1[slen-1-i] {
			return false
		}
	}
	return true



}

func isPalindromeLen(s1 string, l int) bool {

	//fmt.Println("isPalindromelen :",s1,l)	
	for i := 0; i <= l/2; i ++ {
		if s1[i] != s1[l-1-i] {
			return false
		}
	}
	//fmt.Println("isPalindromelen, found Palindrome",s1,l,s1[:l]) 
	return true

}

func findLongestPalindrome(s1 string) (int, int) {
//Given any string, find longest palindrome within that string

	slen := len(s1)
	retLen, idx := -1,0
	//*sPtr := &s1
	for i := 0; i < slen; i ++ {
		for j := i + 1; j < slen; j ++ {
			if isPalindromeLen(s1[i:slen],j-i+1) {
				// found a palindrome
				if  j-i > retLen {
					retLen = j-i
					idx = i
					fmt.Println("palindrome", s1[i:j])
				}
			}
		}
	}
	return retLen, idx
}

func main() {

	/*
	idx := findSubstr(os.Args[1], os.Args[2])
	if idx == -1 {
		fmt.Println("no match found")
	} else {
		fmt.Println("index is at", idx)
	}
	
	wasFound := isPalindrome(os.Args[1])
	if wasFound == true {
		fmt.Println("Palindrome", os.Args[1])
	} else {

		fmt.Println("Not Palindrome", os.Args[1])

	}
	*/
	palLen, startIdx := findLongestPalindrome(os.Args[1])
	fmt.Println("len, index", palLen, startIdx) 
}
