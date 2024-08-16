package reconciliations

import (
	"context"

	"github.com/martinyonatann/acte/internal/reconciliations/dtos"
)

type ReconciliationUC interface {
	TransactionReconciliation(ctx context.Context, request dtos.ReconcileRequest) (dtos.ReconcileResponse, error)
}
