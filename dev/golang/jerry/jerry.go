/*

Given a range of salaries, find percentiles.

 1. Use information from a static array
 	Calculate percentiles

 2. Use information from a CSV file


 Show the following

 - Arithmetic Mean, Geometric Mean, Standard Deviation for a value.
 	value could be base, stock or total

Packages derived from 

- https://godoc.org/golang.org/x/perf/internal/stats

*/


package main

import (
	"fmt"
	"stats" // use this for now, write own package
	"os"
	"encoding/csv"
	"strconv"
)

type Comp struct {

	Level 	int
	YaC	float64 // Years at Company
	YoE	float64 // Years of Experience
	Base	float64
	Stock	float64
	Bonus	float64
	Total	float64
	Started	string // Starting date
}

type Employer struct {

	Name string
	Data []Comp


}

/*
func (emp *Employer) SortData(crit int) Employer {

// Sorts data based on criteria	 and returns a new object
// 0 for date, 1 for total comp

}
*/

func (emp *Employer)  CalcStats (tgt float64, lev int, minYoE float64, maxYoE float64) (float64 , float64) {
// function returns annual base, 4-year RSU offering
// input is target percentile, Level, min. YOE, max. YOE. -1 for no preference
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Printf("|| Records: %d | Level: %d | Percentile: %.2f | %.1f < Exp <  %.1f:     ||\n", len(emp.Data), lev, tgt * 100, minYoE, maxYoE)
	fmt.Println("--------------------------------------------------------------------------")
	var vBase, vStock, vTotal = stats.Sample{Sorted: false}, stats.Sample{Sorted: false}, stats.Sample{Sorted: false}
	for _, val := range emp.Data {
		if  val.Level == lev && val.YoE >= minYoE && val.YoE <= maxYoE { // for a level within an experience range
			//vTotal.Xs = append(vTotal.Xs, (float64) (val.Base + val.Stock + val.Bonus))
			tempTotal := (1.20 * val.Base) + val.Stock // Normalizing bonus as 0.20 of base
			vTotal.Xs = append(vTotal.Xs, tempTotal) 
			vBase.Xs = append(vBase.Xs, val.Base)
			vStock.Xs = append(vStock.Xs, val.Stock)
		}
	}
	fmt.Printf("|| Target Base: %.2f ||\n|| Target Stock: %.2f ||\n|| Target Comp: %.2f ||\n", vBase.Percentile(tgt), vStock.Percentile(tgt), vTotal.Percentile(tgt))
	fmt.Println("------------------------------------------------")
	return vBase.Percentile(tgt), 4*vStock.Percentile(tgt)

}

func (emp *Employer) ReadFromFile(empname string, fname string) {

	// Read from a file and create new entries
	// List could be empty or full
	emp.Name = empname
	fd, err := os.Open(fname)
	if err != nil {
		fmt.Println("file open error", err)
		os.Exit(1)
	}
	defer fd.Close()
	
	vCsvRdr := csv.NewReader(fd)
	rcd, err := vCsvRdr.ReadAll()
	if err != nil  {
		fmt.Println("Reader error : ", err)
		os.Exit(1)
	}
	// fmt.Println("Length of Records : ",len(rcd))
	for _, val := range rcd {
		//fmt.Printf("Record # %d. %s\n", idx, val)
		x := Comp{}
		//TODO: Brace yourself, no error checks here. Will do later 
		x.Level, err = strconv.Atoi(val[0])
		x.YaC, err = strconv.ParseFloat(val[1],64)
		x.YoE, err = strconv.ParseFloat(val[2], 64)
		x.Base, err = strconv.ParseFloat(val[3],64)
		x.Stock, err = strconv.ParseFloat(val[4],64)
		x.Bonus, err = strconv.ParseFloat(val[5],64)
		x.Started = val[7]
		emp.Data = append(emp.Data, x)
	}

}

func main() {

	fb := Employer { Name: "Facebook", Data: []Comp{
		 {
			 6,0,20,225,125,45,0,"6/25/2020",
		 },
		 {
			 6,4,15,220,200,40,0,"2/4/2019",
		 },
		 {
			 6,0,20,210,225,42,0,"1/5/2018",
		 },
		 {
			 6,1,15,220,225,44,0,"9/26/2019",
		 },
		 {
			 6,1,15,220,225,45,0,"7/8/2019",
		 },
		 {
			 6,2,15,225,225,45,0,"5/18/2020",
		 },
		 {
			 6,2.5,20,230,225,45,0,"3/14/2019",
		 },
		 {
			 6,0,15,220,200,88,0,"6/8/2020",
		 },
		 {
			 6,2,20,220,250,44,0,"10/19/2018",
		 },
		 {
			 6,1,20,225,250,45,0,"4/4/2019",
		 },
		 {
			 6,1,18,225,250,50,0,"2/22/2019",
		 },
		 {
			 6,0,20,230,250,46,0,"4/29/2020",
		 },
		 {
			 6,2,18,215,270,50,0,"4/24/2018",
		 },
		 {
			 6,2,15,198,325,40,0,"12/7/2018",
		 },
		 {
			 6,0,15,240,300,48,0,"10/2/2019",
		 },
		 {
			 6,3,15,225,369,43,0,"6/3/2020",
		 },
		 {
			 6,3,15,220,400,40,0,"6/17/2019",
		 },
		 {
			 6,1,20,222,432,39,0,"1/28/2019",
		 },
		 {
			 6,3,16,230,424,46,0,"10/17/2019",
		 },
		 {
			 6,0,20,250,600,50,0,"4/30/2020",
		 },
	     },
	}
	//fmt.Println("FB stats: ", fb)
	vB, vS := fb.CalcStats(0.50,6,15,20)
	fmt.Printf("Annual Offer: Base %.2f & Stock %.2f: \n",vB, vS)
	fmt.Println("*************************************************")

	vB, vS = fb.CalcStats(0.75,6,15,20)
	fmt.Printf("Annual Offer: Base %.2f & Stock %.2f: \n",vB, vS)
	fmt.Println("*************************************************")

	vB, vS = fb.CalcStats(0.80,6,15,20)
	fmt.Printf("Annual Offer: Base %.2f & Stock %.2f: \n",vB, vS)
	fmt.Println("*************************************************")

	fb7 := Employer{}
	fb7.ReadFromFile("Facebook","fb7.csv")
	vB, vS = fb7.CalcStats(0.10,7,15,20)
	fmt.Printf("Annual Offer: Base %.2f & Stock %.2f: \n",vB, vS)
	fmt.Println("*************************************************")


}
