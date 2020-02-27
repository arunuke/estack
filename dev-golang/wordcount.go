// read input and determine count for each individual word 



package main

import (
	"fmt"
	"os"
	"bufio"
)


func main() {


	// Take text file as input

	// Program invoked with one argument
	if (len(os.Args) != 2) {
		fmt.Println("Need exactly one argument. <prog> <file>, os.Args length is -> ",len(os.Args))
		os.Exit(404)
	}

	// Check if file exists and open it
	fd, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File opening error",err)
		os.Exit(404)
	}


	// Create a new	scanner to read the file
	scanner := bufio.NewScanner(fd)

	// Scan the file using words
	scanner.Split(bufio.ScanWords)

	// Declare a map with just one value
	 wordmap := make(map [string] int)
	 

	// Walk through file word by word.
	for scanner.Scan() {
		//If word exists in map, add to count	
		/*
		TODO : This function is case sensitive. Possibly change all words to lower case before calling 
		the comparisons 
		*/
		keystring := scanner.Text()
		if wordmap[keystring] != 0 { // A key with that word exists
			val := wordmap[keystring] // Retrieve the value for that key
			val ++
			wordmap[keystring] = val //Increment value for that key by 1
		} else { // no key with said word exists
			wordmap[keystring] = 1 //Initialize that key with 1 count

		}


		//If word does not exist in map, append to map and initialize


	}

	// Store each word in map as a key-value pair with word and count
	fmt.Println(wordmap)
	
	fd.Close()




}



