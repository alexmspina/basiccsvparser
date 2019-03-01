package basiccsvparser

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"
)

var records = `test,type,thing
1,dog,collar
2,cat,scarf
3,person,wallet
`
var r = csv.NewReader(strings.NewReader(records))

var testHeader = []string{"test", "type", "thing"}

func TestGetHeader(t *testing.T) {
	head := GetHeader(r)
	err := errors.New("result header does not match expected test header")
	if !CompareSlice(head, testHeader) {
		log.Fatal(err)
	}
	fmt.Println("GetHeader test successful")
}
