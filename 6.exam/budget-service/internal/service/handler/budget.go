package handler

import (
	"budget-service/internal/kafka"
	bpb "budget-service/proto/budget"
	"budget-service/proto/transaction"
	"budget-service/storage"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type BudgetService struct {
	bpb.UnimplementedBudgetServiceServer
	DB     *storage.Queries
	Client transaction.ReportServiceClient
	Kafka  *kafka.Kafka
}

func NewBudgetService(db *storage.Queries, cl transaction.ReportServiceClient, k *kafka.Kafka) *BudgetService {
	return &BudgetService{DB: db, Client: cl, Kafka: k}
}

func (b *BudgetService) CreateBudget(ctx context.Context, req *bpb.BudgetInfo) (*bpb.BudgetID, error) {
	user_id := ctx.Value("user_id").(string)
	remIncome, err := b.Client.GetReport(ctx, &transaction.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get response from transaction-microservice: %v", err)
	}
	var message = "budget amount exceeding than income"
	if req.Amount > remIncome.NetSavings {
		if err := b.Kafka.ExceedingBudgetLimit(message, user_id, remIncome); err != nil {
			log.Println(err)
		}
		return nil, fmt.Errorf("budget amount exceeding than income")
	}

	budget_id := uuid.New().String()
	id, err := b.DB.CreateBudget(ctx, storage.CreateBudgetParams{
		UserID:   user_id,
		ID:       budget_id,
		Category: req.Category,
		Amount:   req.Amount,
		Spent:    0,
		Currency: req.Currency,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create budget in postgres: %v", err)
	}
	return &bpb.BudgetID{Id: id}, nil
}

func (b *BudgetService) UpdateBudgetAmount(ctx context.Context, req *bpb.BudgetUpdate) (*bpb.BudgetResponse, error) {
	user_id := ctx.Value("user_id").(string)
	spentAmount, err := b.DB.GetSpentFromBudget(ctx, storage.GetSpentFromBudgetParams{
		UserID: user_id,
		ID:     req.Id,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get spent amount from budget in postgres: %v", err)
	}
	if req.Amount < spentAmount {
		return nil, fmt.Errorf("budget amount cannot be less than spent amount")
	}

	err = b.DB.UpdateBudgetByID(ctx, storage.UpdateBudgetByIDParams{
		UserID: user_id,
		ID:     req.Id,
		Amount: req.Amount,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to update budget in postgres: %v", err)
	}
	return &bpb.BudgetResponse{Message: "Budget Updated Succesfully"}, nil
}
func (b *BudgetService) DeleteBudget(ctx context.Context, req *bpb.BudgetID) (*bpb.BudgetResponse, error) {
	user_id := ctx.Value("user_id").(string)
	err := b.DB.DeleteBudgetByID(ctx, storage.DeleteBudgetByIDParams{
		UserID: user_id,
		ID:     req.Id,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to delete budget in postgres: %v", err)
	}
	return &bpb.BudgetResponse{Message: "Budget Deleted Succesfully"}, nil
}

func (b *BudgetService) GetBudgets(req *bpb.Empty, stream bpb.BudgetService_GetBudgetsServer) error {
	user_id := stream.Context().Value("user_id").(string)
	budgets, err := b.DB.GetBudgets(stream.Context(), user_id)
	if err != nil {
		return fmt.Errorf("unable to get budgets from postgres: %v", err)
	}
	for _, budget := range budgets {
		if err := stream.Send(&bpb.BudgetWithID{
			Id:       budget.ID,
			Category: budget.Category,
			Amount:   budget.Amount,
			Spent:    budget.Spent,
			Currency: budget.Currency,
		}); err != nil {
			return fmt.Errorf("unable to send budget to stream: %v", err)
		}
	}
	return nil
}

func (b *BudgetService) GetBudgetReports(ctx context.Context, req *bpb.Empty) (*bpb.BudgetReport, error) {
	user_id := ctx.Value("user_id").(string)
	report, err := b.DB.GetBudgetReport(ctx, user_id)
	if err != nil {
		return nil, fmt.Errorf("unable to get budget reports  from postgres: %v", err)
	}
	return &bpb.BudgetReport{
		TotalBudget:    float64(report.TotalBudget),
		TotalSpent:     float64(report.TotalSpent),
		RemainingBudget: float64(report.TotalBudget - report.TotalSpent),
	}, nil
}

func (b *BudgetService) AddSpentAmount(ctx context.Context, req *bpb.SpentRequest) (*bpb.BudgetResponse, error) {
	err := b.DB.AddSpentAmount(ctx, storage.AddSpentAmountParams{
		ID:     req.Id,
		Spent: req.Amount,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to add spent amount to budget in postgres: %v", err)
	}
	return &bpb.BudgetResponse{Message: "Spent Amount Added Succesfully"}, nil
}