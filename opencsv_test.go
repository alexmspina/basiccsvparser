package basiccsvparser

import (
	"encoding/csv"
	"fmt"
	"testing"
)

var tester csv.Reader
var filename = "test.csv"

func TestOpenCSV(t *testing.T) {
	testReader := OpenCSV(filename)
	fmt.Println(testReader)
	fmt.Println("OpenCSV test successful")
}
