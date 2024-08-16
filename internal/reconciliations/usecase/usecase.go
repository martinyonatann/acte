package usecase

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/martinyonatann/acte/internal/reconciliations"
	"github.com/martinyonatann/acte/internal/reconciliations/dtos"
	"github.com/martinyonatann/acte/pkg/utils"
)

type reconciliationUC struct {
}

func NewreconciliationUC() reconciliations.ReconciliationUC {
	return &reconciliationUC{}
}

func (uc *reconciliationUC) TransactionReconciliation(ctx context.Context, request dtos.ReconcileRequest) (response dtos.ReconcileResponse, err error) {
	newStartDate, err := time.Parse(time.DateOnly, request.StartDate)
	if err != nil {
		return response, err
	}

	newEndDate, err := time.Parse(time.DateOnly, request.EndDate)
	if err != nil {
		return response, err
	}

	transactionRecords, err := utils.ReadCSV[dtos.TransactionData](request.TransactionPath)
	if err != nil {
		return response, fmt.Errorf("failed to read transaction CSV: %w", err)
	}

	// filter transations with start and end date
	for idx := len(transactionRecords) - 1; idx >= 0; idx-- {
		transaction := transactionRecords[idx]
		if transaction == nil {
			continue
		}

		if !(transaction.TransactionTime.After(newStartDate) && transaction.TransactionTime.Before(newEndDate)) {
			transactionRecords = append(transactionRecords[:idx], transactionRecords[idx+1:]...)
		}
	}

	var savedUnmatchedTransaction = make(map[string]dtos.TransactionData)

	// Process multiple bank statement files
	bankStatementPaths := strings.Split(request.BankStatementPaths, ",")
	for _, bankStatementPath := range bankStatementPaths {
		unmatchedStatements := make(map[string]dtos.BankStatements)
		bankName := strings.Split(filepath.Base(filepath.Dir(bankStatementPath)), "/")[0]

		bankStatementRecords, err := utils.ReadCSV[dtos.BankStatement](bankStatementPath)
		if err != nil {
			return response, fmt.Errorf("failed to read bank %s statement CSV: %w", bankName, err)
		}

		// filter statements with start and end date
		bankStatementMap := make(map[string][]dtos.BankStatement)
		for _, statement := range bankStatementRecords {
			newStatementDate, err := time.Parse(time.DateOnly, statement.Date)
			if err != nil {
				return response, fmt.Errorf("invalid bank %s statement date format in %s", bankName, statement.UniqueIdentifier)
			}

			if newStatementDate.After(newStartDate) && newStatementDate.Before(newEndDate) {
				bankStatementMap[statement.Date] = append(bankStatementMap[statement.Date], *statement)
			}
		}

		// reconcile between transaction and {bank} statements
		for _, transaction := range transactionRecords {
			key := transaction.TransactionTime.Format(time.DateOnly)
			if _, exist := bankStatementMap[key]; exist {
				response.TotalMatched++
				delete(bankStatementMap, key)
			} else {
				if _, exist := savedUnmatchedTransaction[transaction.TransactionID]; !exist {
					response.TotalUnmatched++
					savedUnmatchedTransaction[transaction.TransactionID] = *transaction
				}
			}
		}

		// Handle unmatched bank statements
		for _, statements := range bankStatementMap {
			for _, statement := range statements {
				response.TotalUnmatched++
				unmatchedStatements[bankName] = append(unmatchedStatements[bankName], statement)
			}
		}

		response.UnmatchedStatements = append(response.UnmatchedStatements, unmatchedStatements)
	}

	// handle unmatched transactions
	for _, transaction := range savedUnmatchedTransaction {
		response.UnmatchedTransactions = append(response.UnmatchedTransactions, transaction)
		response.TotalDiscrepancies += transaction.Amount
	}
	response.TotalProcessed = len(transactionRecords)

	return response, nil
}
