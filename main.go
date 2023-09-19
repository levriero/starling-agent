package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
	if len(os.Args) <= 1 {
		log.Fatal("Must pass a CSV file path")
	}

	input, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	inputFileName := input.Name()
	outputFileName := fmt.Sprintf("%v-result.csv", strings.TrimSuffix(inputFileName, filepath.Ext(inputFileName)))

	output, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	defer output.Close()

	generateCSV(input, output)
}

func generateCSV(input, output *os.File) {
	reader := csv.NewReader(bufio.NewReader(input))
	writer := csv.NewWriter(bufio.NewWriter(output))

	defer writer.Flush()

	// Ignore CSV header
	_, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	for {
		row, err := reader.Read()
		if err == io.EOF || err != nil {
			break
		}

		// Convert amount and balance string values to floats
		rowAmount, err := strconv.ParseFloat(row[amount], 64)
		if err != nil {
			log.Fatal(err)
		}
		rowBalance, err := strconv.ParseFloat(row[balance], 64)
		if err != nil {
			log.Fatal(err)
		}

		transaction := Transaction{
			Date:             row[date],
			CounterParty:     row[counterParty],
			Reference:        row[reference],
			Kind:             row[kind],
			SpendingCategory: row[spendingCategory],
			Notes:            row[notes],
			Amount:           rowAmount,
			Balance:          rowBalance,
		}

		notes := buildNotes(transaction)
		csvRow := []string{transaction.Date, fmt.Sprint(transaction.Amount), notes}

		writer.Write(csvRow)
	}

	log.Printf("Statement suceessfully extracted.")
}

// FreeAgent's CSV imports are quite basic, this function retains the data from Starling
// by adding it as a single "note" in FreeAgent's accepted CSV format
// See: https://support.freeagent.com/hc/en-gb/articles/115001222564
func buildNotes(transaction Transaction) string {
	transactionNotes := []string{
		transaction.CounterParty,
		transaction.Reference,
		transaction.Kind,
		transaction.SpendingCategory,
		transaction.Notes,
	}

	return strings.Join(transactionNotes, " ")
}
