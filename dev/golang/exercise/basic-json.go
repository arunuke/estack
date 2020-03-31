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
	Stats		[] GameLog
	Custom		map[string] int
}

func InitPlayer (

func main() {

/*
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
	var PlayerSlice [5]Player

	PlayerSlice[0].Name = "Manning"
	PlayerSlice[0].Position = "QB"
	PlayerSlice[0].Stats[0].GameId = 1
	PlayerSlice[0].Stats[0].PassYd = 200
	PlayerSlice[0].Stats[0].RushYd = 12
	PlayerSlice[0].Stats[0].PassTd = 2
	PlayerSlice[0].Stats[0].RushTd = 0

	fmt.Println("Player Slice :",PlayerSlice)
*/
	DynamicPlayerSlice := make([]Player, 3)
	DynamicPlayerSlice[0].Name = "Manning"
	DynamicPlayerSlice[0].Position = "QB"

	DynamicPlayerSlice[0].Stats[0].GameId = 1
	DynamicPlayerSlice[0].Stats[0].PassYd = 200
	DynamicPlayerSlice[0].Stats[0].RushYd = 12
	DynamicPlayerSlice[0].Stats[0].PassTd = 2
	DynamicPlayerSlice[0].Stats[0].RushTd = 0

	DynamicPlayerSlice[0].Custom = make(map[string] int)
	DynamicPlayerSlice[0].Custom["Int"] = 1
	DynamicPlayerSlice[0].Custom["Sack"] = 2

	fmt.Println("Dynamic Player Slice :",DynamicPlayerSlice)


}
