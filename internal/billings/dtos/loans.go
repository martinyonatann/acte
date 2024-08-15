package dtos

import (
	"time"

	"github.com/invopop/validation"
)

type Loan struct {
	ID                  int64      `json:"id"`
	Amount              float64    `json:"amount"`
	Outstanding         float64    `json:"outstanding"`
	InterestRate        float64    `json:"interest_rate"`
	TermOfPaymentInWeek int        `json:"term_of_payment_in_week"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
}

type CreateLoanRequest struct {
	Amount              float64 `json:"amount"`
	TermOfPaymentInWeek int     `json:"term_of_payment_in_week"`
}

func (c CreateLoanRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Amount, validation.Required),
		validation.Field(&c.TermOfPaymentInWeek, validation.Required),
	)
}

type DetailLoan struct {
	ID                  int64      `json:"id"`
	Amount              float64    `json:"amount"`
	Outstanding         float64    `json:"outstanding"`
	InterestRate        float64    `json:"interest_rate"`
	TermOfPaymentInWeek int        `json:"term_of_payment_in_week"`
	IsDelinquent        bool       `json:"is_delinquent"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`

	LoanSchedules LoanSchedules `json:"schedules"`
}

type LoanRepaymentRequest struct {
	ID          int64   `param:"id"`
	PaymentDate string  `json:"payment_date"`
	Amount      float64 `json:"amount"`
}

func (l LoanRepaymentRequest) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.ID, validation.Required),
		validation.Field(&l.PaymentDate, validation.Required),
		validation.Field(&l.Amount, validation.Required),
	)
}

type InquireRepaymentAmountRequest struct {
	ID          int64  `param:"id"`
	PaymentDate string `json:"payment_date"`
}

type InquireRepaymentAmountResponse struct {
	ID                   int64   `json:"id"`
	PaymentDate          string  `json:"payment_date"`
	TotalRepaymentAmount float64 `json:"total_repayment_amount"`
}

func (l InquireRepaymentAmountRequest) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.ID, validation.Required),
		validation.Field(&l.PaymentDate, validation.Required),
	)
}

type DelinquentLoan struct {
	UnpaidWeeks int `json:"unpaid_weeks"`
	Loan
}

type DelinquentLoans []DelinquentLoan
