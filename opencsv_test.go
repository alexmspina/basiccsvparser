package basiccsvparser

import (
	"encoding/csv"
	"fmt"
	"testing"
)

var tester csv.Reader
var filename = "test.csv"

func TestOpenCSV(t *testing.T) {
	OpenCSV(filename, &tester)
	fmt.Println("OpenCSV test successful")
}
