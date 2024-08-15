package usecase

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/martinyonatann/acte/internal/billings"
	"github.com/martinyonatann/acte/internal/billings/dtos"
	"github.com/martinyonatann/acte/internal/billings/entities"
	"github.com/martinyonatann/acte/pkg/app_error"
	"github.com/martinyonatann/acte/pkg/pagination"
)

type billingUC struct {
	repo billings.BillingRepository
}

func NewBillingUC(repo billings.BillingRepository) billings.BillingUC {
	return &billingUC{repo: repo}
}

func (uc *billingUC) CreateLoan(ctx context.Context, request dtos.CreateLoanRequest) error {
	return uc.repo.Atomic(ctx, &sql.TxOptions{}, func(tx billings.BillingRepository) error {
		loanData := entities.NewLoan(request)
		if err := tx.CreateLoan(ctx, loanData); err != nil {
			return err
		}

		return tx.CreateLoanSchedules(ctx, entities.NewLoanSchedules(request, *loanData))
	})
}

func (uc *billingUC) DetailLoan(ctx context.Context, loanID int64) (loan dtos.DetailLoan, err error) {
	loanData, err := uc.repo.DetailLoan(ctx, loanID)
	if err != nil {
		return loan, err
	}

	return entities.MapDetailLoan(loanData), nil
}

func (uc *billingUC) InquireRepaymentAmount(ctx context.Context, request dtos.InquireRepaymentAmountRequest) (inquiryRepayment dtos.InquireRepaymentAmountResponse, err error) {
	repaymentData, err := uc.repo.GetRepaymentAmount(ctx, request.ID, request.PaymentDate)
	if err != nil {
		return inquiryRepayment, err
	}

	if repaymentData.TotalRepaymentAmount == 0 {
		return inquiryRepayment, app_error.BadRequest(errors.New("payment already made for this date"))
	}

	inquiryRepayment.ID = request.ID
	inquiryRepayment.PaymentDate = request.PaymentDate
	inquiryRepayment.TotalRepaymentAmount = repaymentData.TotalRepaymentAmount

	return inquiryRepayment, nil
}

func (uc *billingUC) LoanRepayment(ctx context.Context, request dtos.LoanRepaymentRequest) error {
	loanData, err := uc.repo.DetailLoan(ctx, request.ID)
	if err != nil {
		return err
	}

	repaymentData, err := uc.repo.GetRepaymentAmount(ctx, loanData.ID, request.PaymentDate)
	if err != nil {
		return err
	}

	if request.Amount != repaymentData.TotalRepaymentAmount {
		return app_error.BadRequest(errors.New("invalid repayment amount"))
	}

	listScheduleNotPaid, err := uc.repo.ListLoanScheduleNotPaid(ctx, loanData.ID, request.PaymentDate)
	if err != nil {
		return err
	}

	return uc.repo.Atomic(ctx, &sql.TxOptions{}, func(tx billings.BillingRepository) error {
		var isPaid = true
		for _, schedule := range *listScheduleNotPaid {
			if err := tx.UpdateLoanSchedule(ctx, entities.UpdateLoanSchedule{ID: schedule.ID, IsPaid: &isPaid, UpdatedAt: time.Now()}); err != nil {
				return err
			}
		}

		newOutstanding := loanData.Outstanding - request.Amount
		return tx.UpdateLoan(ctx, entities.UpdateLoan{ID: loanData.ID, Outstanding: &newOutstanding, UpdatedAt: time.Now()})
	})
}

func (uc *billingUC) ListDelinquentLoan(ctx context.Context, params *pagination.PaginationRequest) (loans dtos.DelinquentLoans, err error) {
	listLoan, err := uc.repo.ListDelinquentLoan(ctx, params)
	if err != nil {
		return loans, err
	}

	return entities.MapDelinquentLoans(listLoan), nil
}
