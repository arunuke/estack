// JSON validations

package main

import (
	"fmt"
	//"os"
	//"encoding/json"
)

// Define regular struct

type RegStr struct {

	Qb string
	Yards int
}

// Define regular struct with exports
type ExpStr struct {
	Qb string `json:"qb"`
	Yards int `json:"yards"`

}

// Define regular struct with maps

type MapStr struct {
	Qb string 
	Stats map [string]int
}

func main() {

	vPlayer1 := RegStr {}
	vPlayer2 := ExpStr { Qb: "Manning", Yards: 540 }
	fmt.Println("P1 : ", vPlayer1)
	fmt.Println("P2 : ", vPlayer2)

	Temp := make(map[string]int)
	Temp["Game1"] = 300
	Temp["Game2"] = 400

	vPlayer3 := MapStr { Qb: "Brady", Stats: Temp }
	vPlayer4 := MapStr { Qb: "Mahomes" }
	fmt.Println("P3 : ", vPlayer3)
	fmt.Println("P4 : ", vPlayer4)



}
