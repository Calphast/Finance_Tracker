package tools

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func AddToCsv(filename string, amount float64, action string) {
	currentTime := time.Now()
	newRow := []string{currentTime.Format("20060102150405"), strconv.FormatFloat(amount, 'E', -1, 64), action}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Write(newRow)
	w.Flush()
}
