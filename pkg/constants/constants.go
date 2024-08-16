package constants

const (
	LoansTableName        = "loans"
	LoanScheduleTableName = "loan_schedules"
)

type TransactionType string

const (
	DebitTransactionType  TransactionType = "DEBIT"
	CreditTransactionType TransactionType = "CREDIT"
)
