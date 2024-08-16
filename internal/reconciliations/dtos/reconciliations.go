package dtos

import (
	"time"

	"github.com/invopop/validation"
	"github.com/martinyonatann/acte/pkg/constants"
)

type ReconcileRequest struct {
	TransactionPath    string `query:"transaction_path"`
	BankStatementPaths string `query:"bank_statement_paths"`
	StartDate          string `query:"start_date"`
	EndDate            string `query:"end_date"`
}

type ReconcileResponse struct {
	TotalProcessed        int                         `json:"total_processed"`
	TotalMatched          int                         `json:"total_matched"`
	TotalUnmatched        int                         `json:"total_unmatched"`
	UnmatchedStatements   []map[string]BankStatements `json:"unmatched_statements"` // Grouped by bank name
	UnmatchedTransactions ListTransactionData         `json:"unmatched_transactions"`
	TotalDiscrepancies    float64                     `json:"total_discrepancies"`
}

func (r ReconcileRequest) Validate() error {

	return validation.ValidateStruct(&r,
		validation.Field(&r.TransactionPath, validation.Required),
		validation.Field(&r.BankStatementPaths, validation.Required),
		validation.Field(&r.StartDate, validation.Required, validation.Date(time.DateOnly)),
		validation.Field(&r.EndDate, validation.Required, validation.Date(time.DateOnly)),
	)
}

type BankStatement struct {
	UniqueIdentifier string  `csv:"unique_identifier" json:"unique_identifier"`
	Amount           float64 `csv:"amount" json:"amount"`
	Date             string  `csv:"date" json:"date"`
}

type TransactionData struct {
	TransactionID   string                    `csv:"trxID" json:"trxID"`
	Amount          float64                   `csv:"amount" json:"amount"`
	Type            constants.TransactionType `csv:"type" json:"type"`
	TransactionTime time.Time                 `csv:"transactionTime" json:"transactionTime"`
}

type BankStatements []BankStatement

type ListTransactionData []TransactionData
