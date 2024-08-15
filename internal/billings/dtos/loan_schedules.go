package dtos

import (
	"time"
)

type LoanSchedule struct {
	ID           int64      `json:"id"`
	LoanID       int64      `json:"loan_id"`
	Amount       float64    `json:"amount"`
	ScheduleDate string     `json:"schedule_date"`
	IsPaid       bool       `json:"is_paid"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type LoanSchedules []LoanSchedule
