package billings

import (
	"context"

	"github.com/martinyonatann/acte/internal/billings/dtos"
	"github.com/martinyonatann/acte/pkg/pagination"
)

type BillingUC interface {
	CreateLoan(ctx context.Context, request dtos.CreateLoanRequest) error
	DetailLoan(ctx context.Context, loanID int64) (loan dtos.DetailLoan, err error)
	LoanRepayment(ctx context.Context, request dtos.LoanRepaymentRequest) error
	InquireRepaymentAmount(ctx context.Context, request dtos.InquireRepaymentAmountRequest) (inquiryRepayment dtos.InquireRepaymentAmountResponse, err error)
	ListDelinquentLoan(ctx context.Context, params *pagination.PaginationRequest) (loans dtos.DelinquentLoans, err error)
}
