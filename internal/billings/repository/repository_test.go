package repository_test

import (
	"context"
	"testing"

	"github.com/martinyonatann/acte/config"
	"github.com/martinyonatann/acte/internal/billings/dtos"
	"github.com/martinyonatann/acte/internal/billings/entities"
	"github.com/martinyonatann/acte/internal/billings/repository"
	"github.com/martinyonatann/acte/pkg/datasources"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var cfg config.Config
var db *gorm.DB

func init() {
	cfg = lo.Must(config.LoadConfigPath("../../../config/config.test"))
	db = lo.Must(datasources.NewMySQLDB(context.Background(), cfg.Database))

}

func Test_Loan(t *testing.T) {
	repo := repository.NewBillingRepository(db)

	t.Run("success create loan", func(t *testing.T) {
		args := dtos.CreateLoanRequest{
			Amount:              5_000_000,
			TermOfPaymentInWeek: 50,
		}

		newLoan := entities.NewLoan(args)

		err := repo.CreateLoan(context.Background(), newLoan)
		require.NoError(t, err)

		DetailLoan, err := repo.DetailLoan(context.Background(), newLoan.ID)
		require.NoError(t, err)
		require.Equal(t, newLoan.Amount, DetailLoan.Amount)
		require.Equal(t, newLoan.Outstanding, DetailLoan.Outstanding)
		require.Equal(t, newLoan.InterestRate, DetailLoan.InterestRate)
		require.Equal(t, newLoan.TermOfPaymentInWeek, DetailLoan.TermOfPaymentInWeek)

		newSchedules := entities.NewLoanSchedules(args, *newLoan)
		err = repo.CreateLoanSchedules(context.Background(), newSchedules)
		require.NoError(t, err)

		LoanSchedules, err := repo.LoanSchedules(context.Background(), newLoan.ID)
		require.NoError(t, err)
		require.Equal(t, len(*LoanSchedules), len(*newSchedules))
	})
}
