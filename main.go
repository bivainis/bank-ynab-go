package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

// TODO: read configuration from .env
// TODO: tests
func main() {
    exportedCsv, err := os.Open("working-dir/exported-from-bank.csv")
    if err != nil {
        log.Fatal(err)
    }

    // remember to close the file at the end of the program
    defer exportedCsv.Close()

	r := csv.NewReader(exportedCsv)
	r.Comma = ';'

	csvFile, err := os.Create("working-dir/ready-for-ynab.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)

	// Column index map, 0 based
	columnIndex := map[string]int{
		"Date": 3,
		"Payee": 4,
		"Memo": 8,
		"Amount": 5,
	}

	index := 0;

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		date := record[columnIndex["Date"]]
		payee := record[columnIndex["Payee"]]
		amount := record[columnIndex["Amount"]]
		bankMemo := record[columnIndex["Memo"]]
		newMemo := ""

		// Override date with a more precise one recorded in memo
		//   (bank transaction sometimes takes up to a few days to clear, but we want the original date)
		if strings.Contains(bankMemo, "bsdato") {
			newMemo = "Bank date: " + date

			// Find date of payment in the memo from bank
			dateRegex := regexp.MustCompile(`\b(\d{4}-\d{2}-\d{2})\b`)
			date = dateRegex.FindStringSubmatch(bankMemo)[1]

			t, err := time.Parse("2006-01-02", date)
			
			if err != nil {
				fmt.Println(err)
				return
			}

			// Math the format of bank statement date
			date = t.Format("02-01-2006")
		}

		payeeFormatted := strings.ReplaceAll(payee, "ø", "X")
		dateFormatted := strings.ReplaceAll(date, "-", "/")

		// Note: instead of both "Outflow" and "Inflow", we can use "Amount", which
		//   is then put into right column based on positive or negative value.
		toWrite := []string{"Date", "Payee", "Memo", "Amount"}

		if index > 0 {
			toWrite = []string{dateFormatted, payeeFormatted, newMemo, amount}
		}

		if err := w.Write(toWrite); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}

		index++
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
