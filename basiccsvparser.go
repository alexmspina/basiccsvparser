package basiccsvparser

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// CompareSlice checks slices for equality
func CompareSlice(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	} else if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// GetHeader from csv files
func GetHeader(cr *csv.Reader) []string {
	record, err := cr.Read()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}

	return record
}

// OpenCSV opens a csv file
func OpenCSV(f string, r *csv.Reader) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	r = csv.NewReader(file)
}

// ParseCsv parse csv file given file name string and deposit in list of maps for each row
func ParseCsv(header []string, r *csv.Reader) []map[string]string {
	// initialize resulting list of maps representing each row
	result := make([]map[string]string, 0)

	// loop through csv reader line by line
	for {
		// read line of csv reader
		record, err := r.Read()
		// break loop when end of file is reached
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err) // log error if not nil
		} else if CompareSlice(record, header) {
			continue // ignore first line ie csv header
		}

		// initialize a csv header key / row value map
		row := make(map[string]string, 0)

		// for each header key create a new value in the row map
		for i, k := range header {
			if record[i] == "" {
				continue
			} else {
				row[k] = record[i]
			}
		}

		result = append(result, row)
	}

	return result
}
