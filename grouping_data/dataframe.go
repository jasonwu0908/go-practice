package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {
	irisCsv, err := os.Open("data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	df := dataframe.ReadCSV(irisCsv)
	head := df.Subset([]int{0, 3})
	fmt.Println(head)
}
