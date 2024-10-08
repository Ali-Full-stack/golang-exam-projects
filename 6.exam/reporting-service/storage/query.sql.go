// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package storage

import (
	"context"
)

const createReport = `-- name: CreateReport :exec
INSERT INTO reports (id, income, expense, net_saving, budget_amount, budget_spent, remaining_budget)
VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type CreateReportParams struct {
	ID              string
	Income          float64
	Expense         float64
	NetSaving       float64
	BudgetAmount    float64
	BudgetSpent     float64
	RemainingBudget float64
}

func (q *Queries) CreateReport(ctx context.Context, arg CreateReportParams) error {
	_, err := q.db.ExecContext(ctx, createReport,
		arg.ID,
		arg.Income,
		arg.Expense,
		arg.NetSaving,
		arg.BudgetAmount,
		arg.BudgetSpent,
		arg.RemainingBudget,
	)
	return err
}
