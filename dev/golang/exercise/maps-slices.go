// JSON validations

// Create a constructor for the objects
package main

import (
	"fmt"
	//"os"
	//"encoding/json"
)


type GameLog struct {

	GameId 	int
	PassYd	int
	RushYd	int
	PassTd	int
	RushTd 	int

}

type Player struct {
	Name		string
	Position	string	
	//Gameday stats in an array
	Stats		[] GameLog
	//Arbitrary keys and values
	Custom		map[string] int
}

func NewPlayer ( N string, P string ) *Player {

	/*
	Cannot use make for initializing a struct
	'make' can be used only for slices and maps
	Use new for everything else
	*/
	p := new (Player)

	// You got to initialize elements inside the struct
	p.Stats = make([]GameLog,1) 
	p.Custom = make(map[string] int) 

	p.Name = N
	p.Position = P
	return p

}

func BasicTests () {


	// Basic Tests for slices, maps and constructors

	// Create a new object using a constructor
	newp := NewPlayer( "Brady", "QB")
	newp.Custom["Inc"] = 9
	newp.Custom["Hurry"] = 3
	// Printing as a dereferenced pointer
	fmt.Println("New Player", *newp)

	// Create a new object using constructor
	p2 := NewPlayer("Romo", "QB")
	// Printing as a pointer
	fmt.Println("P2", p2)

	// Creating an individual stand-alone object without constructor
	// Use struct literals to initialize
	dp := Player { }
	dp.Name = "Manning"
	dp.Position = "QB"
	/*
	Cannot add directly to "Custom" since it hasn't been initialized
	Ideally, use a constructor for all struct elements
	No need to declare and assign using := since we just initialized
	earlier
	*/
	dp.Custom = make(map[string] int)
	dp.Custom["Inc"] = 7
	fmt.Println("dp", dp)


	// Create slice of players where all individual players will be added
	vPlArr := make ([]Player,1)
	vPlArr[0].Name = "Sherman"
	// Cannot add values to 'Custom' since make initialized only the Player array, 
	// It did not initialize the internal arrays and maps
	fmt.Println("vPlaArr", vPlArr)

	// Add every single player
	vPlArr = append(vPlArr, *newp, *p2, dp)
	fmt.Println("vPlaArr", vPlArr)

}

func main() {

	BasicTests()

}
