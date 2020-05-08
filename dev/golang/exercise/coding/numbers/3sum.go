/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]o
*/

package main

import (
	"fmt"
	"math/rand"
)

func populateArr (a *[]int, s int) {

	for i := 0;  i < s; i ++ {
		*a = append(*a, rand.Int()%10)
	}
	fmt.Println("In populateArr",a) 

}

func arrBubbleSort(a *[]int) {

	slen := len(*a)
	for i := 0; i < slen; i ++ {
		for j := 0; j < slen; j ++ {
			if (*a)[i] > (*a)[j] {
				(*a)[j], (*a)[i] = (*a)[i], (*a)[j]
			}
		}
	}
	fmt.Println("In arrBubbleSort",a) 

}

func calcSum (a []int, s int) {

	i,j,k := 0,0,0
	aLen := len(a)
	for i = 0; i < aLen ; i ++ {
		for j = i +1; j <aLen ; j ++ {
			for k = j +1; k <aLen ; k ++ {
				if  a[i] + a[j] + a[k] == 0 {
					fmt.Printf("\n[ %d, %d, %d ]\n",a[i], a[j], a[k])
				}
			}
		}
	}
}

func calcSumWithHash(a []int, s int) {
// a - array of integers, s - sum to calculate




}


func main() {

	//inputArr := make([]int, 10) 
	var inputArr []int
	maxNum := 10
	finSum := 0
	populateArr(&inputArr, maxNum)
	fmt.Println("In main",inputArr) 
	arrBubbleSort(&inputArr)
	fmt.Println("In main, sorted",inputArr) 

	inputArr2 := []int { -1, 0, 1, 2, -1, -4 }
	calcSum(inputArr2, finSum)



}
