package entities

import (
	"time"

	"github.com/martinyonatann/acte/internal/billings/dtos"
	"github.com/martinyonatann/acte/pkg/constants"
)

type LoanSchedule struct {
	ID           int64      `gorm:"column:id;primaryKey"`
	LoanID       int64      `gorm:"column:loan_id"`
	Amount       float64    `gorm:"column:amount"`
	ScheduleDate string     `gorm:"column:schedule_date"`
	IsPaid       bool       `gorm:"column:is_paid"`
	CreatedAt    time.Time  `gorm:"column:created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at"`
}

type UpdateLoanSchedule struct {
	ID           int64     `gorm:"column:id;primaryKey"`
	LoanID       *int64    `gorm:"column:loan_id"`
	Amount       *float64  `gorm:"column:amount"`
	ScheduleDate *string   `gorm:"column:schedule_date"`
	IsPaid       *bool     `gorm:"column:is_paid"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

type LoanSchedules []LoanSchedule

func (LoanSchedule) TableName() string {
	return constants.LoanScheduleTableName
}

func NewLoanSchedules(request dtos.CreateLoanRequest, loan Loan) *LoanSchedules {
	schedules := make(LoanSchedules, 0)
	currentDate := time.Now()
	scheduleAmount := loan.Outstanding / float64(loan.TermOfPaymentInWeek)

	for i := 1; i <= request.TermOfPaymentInWeek; i++ {
		schedules = append(schedules, LoanSchedule{
			LoanID:       loan.ID,
			Amount:       scheduleAmount,
			ScheduleDate: currentDate.AddDate(0, 0, 7*(i-1)).Format(time.DateOnly),
		})
	}

	return &schedules
}

type GetRepaymentAmount struct {
	ID                   int64   `gorm:"column:id;primaryKey"`
	TotalRepaymentAmount float64 `gorm:"column:total_repayment_amount"`
}
