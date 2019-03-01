package basiccsvparser

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"
)

var parsingRecords = `test,type,thing
1,dog,collar
2,cat,scarf
3,person,wallet
`
var cr = csv.NewReader(strings.NewReader(parsingRecords))

var parseTestHeader = []string{"test", "type", "thing"}
var record1 = map[string]string{
	"test":  "1",
	"type":  "dog",
	"thing": "collar",
}
var record2 = map[string]string{
	"test":  "2",
	"type":  "cat",
	"thing": "scarf",
}
var record3 = map[string]string{
	"test":  "3",
	"type":  "person",
	"thing": "wallet",
}

var testRecords = []map[string]string{record1, record2, record3}

func TestParseCSV(t *testing.T) {
	result := ParseCsv(parseTestHeader, cr)
	err := CheckResult(result)
	if err != nil {
		log.Fatal(err)
	}
	for i, t := range testRecords {
		for k, v := range t {
			if val, ok := result[i][k]; ok {
				if val == v {
					continue
				} else {
					err = errors.New("record " + string(i+1) + " value result does not match expected test value")
					log.Fatal(err)
				}
			} else {
				err = errors.New("record " + string(i+1) + " key result does not match expected test value")
				log.Fatal(err)
			}
		}
	}
	fmt.Println("parseCSV test successful")
}

// CheckReport tests if struct is of type Report
func CheckResult(r interface{}) error {
	switch r.(type) {
	case []map[string]string:
		return nil
	default:
		return errors.New("not a []map[string]string")
	}
}
