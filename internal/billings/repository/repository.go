package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/martinyonatann/acte/internal/billings"
	"github.com/martinyonatann/acte/internal/billings/entities"
	"github.com/martinyonatann/acte/pkg/app_error"
	"github.com/martinyonatann/acte/pkg/constants"
	"github.com/martinyonatann/acte/pkg/pagination"
	"github.com/sourcegraph/conc/pool"
	"gorm.io/gorm"
)

type billingRepo struct {
	db *gorm.DB
}

func NewBillingRepository(db *gorm.DB) billings.BillingRepository {
	return &billingRepo{db: db}
}

func (r *billingRepo) Atomic(ctx context.Context, opt *sql.TxOptions, cb func(tx billings.BillingRepository) error) error {
	tx := r.db.Begin(opt)

	repo := &billingRepo{
		db: tx,
	}

	if err := cb(repo); err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return err
		}
	}

	return tx.Commit().Error
}

func (r *billingRepo) CreateLoan(ctx context.Context, loan *entities.Loan) error {
	return r.db.WithContext(ctx).Table(constants.LoansTableName).Create(loan).Error
}

func (r *billingRepo) DetailLoan(ctx context.Context, loanID int64) (loan entities.DetailLoan, err error) {
	if err := r.db.WithContext(ctx).Table(constants.LoansTableName).Preload("LoanSchedules").Where("id = ?", loanID).Take(&loan).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return loan, app_error.NotFound(errors.New("loan not found"))
		}

		return loan, err
	}

	return loan, nil
}

func (r *billingRepo) UpdateLoan(ctx context.Context, loan entities.UpdateLoan) error {
	return r.db.WithContext(ctx).Table(constants.LoansTableName).Updates(loan).Error
}

func (r *billingRepo) CreateLoanSchedules(ctx context.Context, schedules *entities.LoanSchedules) error {
	return r.db.WithContext(ctx).Table(constants.LoanScheduleTableName).Create(schedules).Error
}

func (r *billingRepo) LoanSchedules(ctx context.Context, loanID int64) (schedules *entities.LoanSchedules, err error) {
	err = r.db.WithContext(ctx).Table(constants.LoanScheduleTableName).Where("loan_id = ?", loanID).Find(&schedules).Error
	if err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *billingRepo) ListLoanScheduleNotPaid(ctx context.Context, loanID int64, repaymentDate string) (schedules *entities.LoanSchedules, err error) {
	err = r.db.WithContext(ctx).Table(constants.LoanScheduleTableName).
		Where("loan_id = ?", loanID).
		Where("is_paid IS NOT TRUE").
		Where("schedule_date <= DATE(?)", repaymentDate).
		Find(&schedules).Error
	if err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *billingRepo) UpdateLoanSchedule(ctx context.Context, request entities.UpdateLoanSchedule) error {
	return r.db.Table(constants.LoanScheduleTableName).Updates(request).Error
}

func (r *billingRepo) GetRepaymentAmount(ctx context.Context, loanID int64, repaymentDate string) (response entities.GetRepaymentAmount, err error) {
	err = r.db.WithContext(ctx).Raw(`
			SELECT
			    l.id,
			    SUM(ls.amount) AS total_repayment_amount
			FROM loans l
			JOIN loan_schedules ls ON ls.loan_id = l.id
			WHERE
			    l.id = ? 
			    AND ls.is_paid IS NOT TRUE
			    AND ls.schedule_date <= DATE(?)
			GROUP BY
			    l.id;`, loanID, repaymentDate).Take(&response).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, app_error.NotFound(errors.New("no repayment scheduled for the provided date"))
		}

		return response, err
	}

	return response, nil
}

func (r *billingRepo) ListDelinquentLoan(ctx context.Context, params *pagination.PaginationRequest) (loans entities.DelinquentLoans, err error) {
	query := r.db.WithContext(ctx).
		Table("loans l").
		Select("l.id, l.amount, l.outstanding, l.term_of_payment_in_week, l.created_at, l.updated_at, COUNT(*) AS unpaid_weeks").
		Joins("JOIN loan_schedules ls ON ls.loan_id = l.id").
		Where("ls.is_paid IS NOT TRUE AND ls.schedule_date <= DATE(?)", time.Now().Format(time.DateOnly)).
		Group("l.id").
		Having("COUNT(*) >= 2")

	workers := pool.New().WithContext(ctx)
	workers.Go(func(ctx context.Context) error {
		return query.Session(&gorm.Session{}).Order(params.GetSort()).Limit(params.GetLimit()).Offset(params.GetOffset()).Find(&loans).Error
	})
	workers.Go(func(ctx context.Context) error {
		return query.Session(&gorm.Session{}).Count(&params.Total).Error
	})

	if workers.Wait() != nil {
		return loans, err
	}

	return loans, nil
}
