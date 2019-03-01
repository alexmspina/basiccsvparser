package basiccsvparser

import (
	"errors"
	"fmt"
	"log"
	"testing"
)

var sliceCompareCases = [][][]string{
	[][]string{
		[]string{"taco", "taco"},
		[]string{"taco", "taco"},
	},
	[][]string{
		[]string{"meatball", "sausage"},
		[]string{"sausage", "meatball"},
	},
	[][]string{
		[]string{"pizza", "Pizza"},
		[]string{"pizza", "pizza"},
	},
	[][]string{
		[]string{"dog", "dog", "cat", "lizard"},
		[]string{"dog", "dog", "cat", "lizard"},
	},
}

var sliceCompareResults = []bool{true, false, false, true}

func TestSliceCompare(t *testing.T) {
	err := errors.New("string compare failed")

	for i, pair := range sliceCompareCases {
		if CompareSlice(pair[0], pair[1]) != sliceCompareResults[i] {
			fmt.Printf("Does slice 1, %v, match slice 2, %v? Result was %v, but should be %v", pair[0], pair[1], sliceCompareResults[i], !sliceCompareResults[i])
			log.Fatal(err)
		} else {
			continue
		}
	}

	fmt.Println("CompareSlice test successful")
}
