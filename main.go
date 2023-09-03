package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Transaction struct {
	Date             string
	CounterParty     string
	Amount           float64
	Balance          float64
	Reference        string
	Kind             string
	SpendingCategory string
	Notes            string
}

const (
	date int = iota
	counterParty
	reference
	kind
	amount
	balance
	spendingCategory
	notes
)

func main() {
	f, err := os.Open("statement.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := csv.NewReader(f)

	// Ignore CSV header
	_, err = reader.Read()

	if err != nil {
		log.Fatal(err)
	}

	for {
		row, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		// Break if we can't parse the transaction amount
		rowAmount, err := strconv.ParseFloat(row[amount], 64)

		if err != nil {
			log.Fatal(err)
		}

		rowBalance, err := strconv.ParseFloat(row[balance], 64)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(Transaction{
			Date:             row[date],
			CounterParty:     row[counterParty],
			Reference:        row[reference],
			Kind:             row[kind],
			SpendingCategory: row[spendingCategory],
			Notes:            row[notes],
			Amount:           rowAmount,
			Balance:          rowBalance,
		})
	}
}
