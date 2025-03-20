package tools

import (
	"encoding/csv"

	"fmt"
	"log"
	"os"
	"time"
)

func AddToCsv(filename string, amount float64, action string) {
	currentTime := time.Now()
	newRow := []string{currentTime.Format("20060102150405"), USD(amount).String(), action}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Write(newRow)
	w.Flush()
}

func ReadCsv(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	record, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range record {
		fmt.Print(row[1] + " " + row[2])
		fmt.Print("\n")
	}

	return record
}
