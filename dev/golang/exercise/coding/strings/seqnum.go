
package main


import (
	"fmt"
	"os"
//	"strconv"
)

func AddSequence(upper int) int {

	sum := 0
	for i:=0;i <= upper; i++ {
		sum += i
	}
	return sum

}

func IsUnique(str []byte) {

	slen := len(str)
	for i:=0;i<slen-1;i++ {
		for j:=i+1; j< slen; j++ {
			if str[i] == str[j] {
				fmt.Printf("Not unique, Twice occuring character is %c\n",str[i])
				return
			}
		}
	}
	fmt.Printf("All characters are unique in %s\n", str)
	return

}

func IsPermutation(str1 []byte, str2 []byte) {


	//Calculate length of strings
	s1len := len(str1)
	s2len := len(str2)

	//If strings are not same length, return right away
	if  s1len != s2len {
		fmt.Println("lengths not equal")
		return
	}

	//Convoluted algorithm
	//Pick a letter in first string, check every character in second string
	//If you have reached end of string, you didn't find a match. so not a permutation
	//End of string is tracked via tracker variable 
	for i := 0; i < s1len; i ++ {
		tracker := s1len 
		for j := 0; j < s2len; j ++  {
			if str1[i] == str2[j] {
				//If a match was found, break away
				break
			}
			tracker -- 
		}
		if tracker == 0 {
			fmt.Printf("%c not found in %s\n",str1[i],str2)
			return
		}
	}
	fmt.Printf("%s is a permutation of %s\n",str1, str2)

}

func PalindromePermutation(str []byte) {

	// Calculate length of string
	slen := len(str)

	// no need for ret. it just became a pointer to the same string.

	fmt.Printf("Input %s\n",str)
	for i := 0; i < slen; i ++ {
		for j := i+1; j < slen; j ++ {
			if str[i] == str[j] {
				//Found a match. swap the new character with
				//the one that fits the other end of the palindrome 
				mir :=  slen - 1 - i
				tch := str[j]
				str [j] = str [mir]
				str [mir] = tch
			}
		}
	}
	fmt.Printf("Palindrome: %s\n",str)

}

func AddCharsInSpace(str []byte) {



}

func main() {

	// arg1 : the top end of the sequence till which to find the sum

	fmt.Println("1. upper bound sum of sequence")
	fmt.Println("2. string for unique characters")
	fmt.Println("3. string1 and string 2 for permutation ")
	fmt.Println("4. adding %20 as a character for space")

	fmt.Println("Args", os.Args[0], os.Args[1])

	/*
	// Problem 1
	upper, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error in strconv",err)
		os.Exit(1)
	} else {
		seqsum := AddSequence(upper)
		fmt.Printf("\nUpper bound %d, Sum %d\n",upper,seqsum);
	}
	*/

	// Problem 2
	// IsUnique([]byte (os.Args[1]))

	// Problem 3
	//IsPermutation([]byte (os.Args[1]), []byte (os.Args[2]))

	// Problem 4
	PalindromePermutation([]byte(os.Args[1]))

}

