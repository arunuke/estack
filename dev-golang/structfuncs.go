/*

Playing around with structures and functions
 Creating nested structures
 Creating methods for structures
 Creating constructors for structures

To DO

Write a New function for the objects
If a struct has a nested struct, how to initialize it in the command line?
Ans: initialize a zero object, then assign values

*/

package main
import (
	"fmt"
	"os"
)


//Define a struct titled Point to hold a location on a graph (x-axis and y-axis)
type Point struct {
	x int // holds x-axis
	y int // holds y-axis

}

//Define a struct titled Square to hold dimensions of a square
type Square struct {

	center Point // holds center of square
	length int // length of a side

}

//Define a function to move a square from it's current position by the new values
func (s *Square) Move (dx int, dy int) {

	s.center.x += dx
	s.center.y += dy

}

//Define a function to calculate area of a square
func (s *Square) Area () (int) {

	area := s.length * s.length
	fmt.Println("Area of square is -->", area)
	return area

}

// New function for Square
func  NewSquare (x int, y int, len int) (*Square, error) {
	if x < 0 || y < 0 || len < 0 {
		return nil, fmt.Errorf("All values need to be above zero")
	}
	s := &Square{}
	s.center.x = x
	s.center.y = y
	s. length = len
	return s, nil	

}

func main() {

	sq, err := NewSquare(3, 3, 4) 
	if err != nil {
		fmt.Println("Newsquare returned error", err)
		os.Exit(404) 
	}
	sq.Move(4,4)
	sq.Area()


}
