package billings

import (
	"context"
	"database/sql"

	"github.com/martinyonatann/acte/internal/billings/entities"
	"github.com/martinyonatann/acte/pkg/pagination"
)

type BillingRepository interface {
	Atomic(ctx context.Context, opt *sql.TxOptions, cb func(tx BillingRepository) error) error

	CreateLoan(ctx context.Context, loan *entities.Loan) error
	DetailLoan(ctx context.Context, loanID int64) (loan entities.DetailLoan, err error)
	UpdateLoan(ctx context.Context, loan entities.UpdateLoan) error
	ListDelinquentLoan(ctx context.Context, params *pagination.PaginationRequest) (loans entities.DelinquentLoans, err error)

	CreateLoanSchedules(ctx context.Context, schedules *entities.LoanSchedules) error
	UpdateLoanSchedule(ctx context.Context, request entities.UpdateLoanSchedule) error
	LoanSchedules(ctx context.Context, loanID int64) (schedules *entities.LoanSchedules, err error)
	ListLoanScheduleNotPaid(ctx context.Context, loanID int64, repaymentDate string) (schedules *entities.LoanSchedules, err error)
	GetRepaymentAmount(ctx context.Context, loanID int64, repaymentDate string) (response entities.GetRepaymentAmount, err error)
}
