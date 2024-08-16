package entities

import (
	"time"

	"github.com/martinyonatann/acte/internal/billings/dtos"
	"github.com/martinyonatann/acte/pkg/constants"
)

const (
	InterestRate float64 = 0.10
)

type Loan struct {
	ID                  int64      `gorm:"column:id;primaryKey"`
	BorrowerID          int64      `gorm:"column:borrower_id"`
	Amount              float64    `gorm:"column:amount"`
	Outstanding         float64    `gorm:"column:outstanding"`
	InterestRate        float64    `gorm:"column:interest_rate"`
	TermOfPaymentInWeek int        `gorm:"column:term_of_payment_in_week"`
	CreatedAt           time.Time  `gorm:"column:created_at"`
	UpdatedAt           *time.Time `gorm:"column:updated_at"`
}

type DelinquentLoan struct {
	UnpaidWeeks int `gorm:"column:unpaid_weeks"`
	Loan
}

type DelinquentLoans []DelinquentLoan

type UpdateLoan struct {
	ID                  int64     `gorm:"column:id;primaryKey"`
	BorrowerID          *int64    `gorm:"column:borrower_id"`
	Amount              *float64  `gorm:"column:amount"`
	Outstanding         *float64  `gorm:"column:outstanding"`
	InterestRate        *float64  `gorm:"column:interest_rate"`
	TermOfPaymentInWeek *int      `gorm:"column:term_of_payment_in_week"`
	UpdatedAt           time.Time `gorm:"column:updated_at"`
}

func (Loan) TableName() string {
	return constants.LoansTableName
}

func NewLoan(request dtos.CreateLoanRequest) *Loan {
	return &Loan{
		BorrowerID:          request.BorrowerID,
		Amount:              request.Amount,
		InterestRate:        InterestRate,
		TermOfPaymentInWeek: request.TermOfPaymentInWeek,
		Outstanding:         request.Amount * (1 + InterestRate),
		CreatedAt:           time.Now(),
	}
}

type DetailLoan struct {
	ID                  int64      `gorm:"column:id;primaryKey"`
	BorrowerID          int64      `gorm:"column:borrower_id"`
	Amount              float64    `gorm:"column:amount"`
	Outstanding         float64    `gorm:"column:outstanding"`
	InterestRate        float64    `gorm:"column:interest_rate"`
	TermOfPaymentInWeek int        `gorm:"column:term_of_payment_in_week"`
	CreatedAt           time.Time  `gorm:"column:created_at"`
	UpdatedAt           *time.Time `gorm:"column:updated_at"`

	LoanSchedules LoanSchedules `gorm:"foreignKey:LoanID"`
}

func (l DetailLoan) IsDelinquent() bool {
	missedPayments := 0

	for _, payment := range l.LoanSchedules {
		if !payment.IsPaid {
			missedPayments++
			if missedPayments >= 2 {
				return true
			}
		} else {
			missedPayments = 0
		}
	}
	return false
}

func MapDetailLoan(l DetailLoan) dtos.DetailLoan {
	var schedules = make(dtos.LoanSchedules, 0)
	for _, schedule := range l.LoanSchedules {
		schedules = append(schedules, dtos.LoanSchedule{
			ID:           schedule.ID,
			LoanID:       schedule.LoanID,
			Amount:       schedule.Amount,
			ScheduleDate: schedule.ScheduleDate,
			IsPaid:       schedule.IsPaid,
			CreatedAt:    schedule.CreatedAt,
			UpdatedAt:    schedule.UpdatedAt,
		})
	}

	return dtos.DetailLoan{
		ID:                  l.ID,
		BorrowerID:          l.BorrowerID,
		Amount:              l.Amount,
		Outstanding:         l.Outstanding,
		InterestRate:        l.InterestRate,
		IsDelinquent:        l.IsDelinquent(),
		TermOfPaymentInWeek: l.TermOfPaymentInWeek,
		CreatedAt:           l.CreatedAt,
		UpdatedAt:           l.UpdatedAt,

		LoanSchedules: schedules,
	}
}

func MapDelinquentLoans(loans DelinquentLoans) dtos.DelinquentLoans {
	var listLoan = make(dtos.DelinquentLoans, 0)
	for _, loan := range loans {
		listLoan = append(listLoan, dtos.DelinquentLoan{
			UnpaidWeeks: loan.UnpaidWeeks,
			Loan: dtos.Loan{
				ID:                  loan.ID,
				BorrowerID:          loan.BorrowerID,
				Amount:              loan.Amount,
				Outstanding:         loan.Outstanding,
				InterestRate:        InterestRate,
				TermOfPaymentInWeek: loan.TermOfPaymentInWeek,
				CreatedAt:           loan.CreatedAt,
				UpdatedAt:           loan.UpdatedAt,
			},
		})
	}

	return listLoan
}
