package handler

import (
	"context"
	"report-service/proto/budget"
	rpb "report-service/proto/report"
	"report-service/proto/transaction"
	"report-service/storage"

	"github.com/google/uuid"
)

type ReportService struct {
	rpb.UnimplementedReportServiceServer
	BudgetClient      budget.BudgetServiceClient
	TransactionClient transaction.TransactionServiceClient
	DB                *storage.Queries
}

func NewReportService(bc budget.BudgetServiceClient, tc transaction.TransactionServiceClient, db *storage.Queries) *ReportService {
	return &ReportService{BudgetClient: bc, TransactionClient: tc, DB: db}
}

func (r *ReportService) GetReport(ctx context.Context, in *rpb.ReportEmpty) (*rpb.ReportResponse, error) {
	budgetReport, err := r.BudgetClient.GetBudgetReports(ctx, &budget.Empty{})
	if err != nil {
		return nil, err
	}
	transactionReport, err := r.TransactionClient.GetReport(ctx, &transaction.Empty{})
	if err != nil {
		return nil, err
	}
	id := uuid.New().String()
	r.DB.CreateReport(ctx, storage.CreateReportParams{
		ID:              id,
		Income:          transactionReport.Income,
		Expense:         transactionReport.Expenses,
		NetSaving:       transactionReport.NetSavings,
		BudgetAmount:    budgetReport.TotalBudget,
		BudgetSpent:     budgetReport.TotalSpent,
		RemainingBudget: budgetReport.RemainingBudget,
	})

	return &rpb.ReportResponse{
		Income: transactionReport.Income,
		Expenses: transactionReport.Expenses,
		NetSavings: transactionReport.NetSavings,
		Budget: &rpb.Budget{
			TotalAmount: budgetReport.TotalBudget,
			TotalSpent: budgetReport.TotalSpent,
			RemainingBudget: budgetReport.RemainingBudget,
		},
	}, nil
}

