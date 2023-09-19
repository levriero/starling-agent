package main

import "testing"

func TestMain(t *testing.T) {

}

func TestGenerateCSV(t *testing.T) {

}

func TestBuildNotes(t *testing.T) {
	transaction := Transaction{
		Date:             "01/02/2023",
		CounterParty:     "ACME CORP",
		Reference:        "Bill Payment",
		Kind:             "TRANSFER",
		SpendingCategory: "",
		Notes:            "",
		Amount:           1250.80,
		Balance:          2000.00,
	}

	got := buildNotes(transaction)
	want := "ACME CORP Bill Payment TRANSFER  "

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
