// All sorting algorithms 

package main
import (
	"fmt"
	"math/rand"
)

type Array struct {

	Data []int // Slice of integers
	Items int // Number of elements

}

func NewArray(count int) *Array {

	retArr := new(Array)
	retArr.Data = make([]int, count)
	retArr.Items = count

	return retArr

}

func (a *Array) Load() {
// Load an array with random integers for future processing

	for i := 0; i < a.Items; i ++ {
		a.Data[i] = rand.Int() % 100
	}


}

func (a *Array) BubbleSort() {
	for i := 0; i < a.Items; i ++ {
		for j := 0; j < a.Items; j ++ {
			if a.Data[i] > a.Data[j] {
				a.Data[i], a.Data[j] = a.Data[j], a.Data[i]
			}
		}
	}

}



func main() {
	max := 10
	vArray := NewArray(max)
	fmt.Println("vArray Initialized. ",vArray.Data, vArray.Items)
	vArray.Load()
	fmt.Println("vArray Loaded. ",vArray.Data, vArray.Items)
	vArray.BubbleSort()
	fmt.Println("vArray BubbleSorted",vArray)

}
