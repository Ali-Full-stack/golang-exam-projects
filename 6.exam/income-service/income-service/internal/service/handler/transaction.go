package handler

import (
	"context"
	"fmt"
	"time"
	"transaction-service/kafka"
	tpb "transaction-service/proto/income"
	"transaction-service/storage"

	"github.com/google/uuid"
)

var (
	MaxIncome  = 100000
	MaxExpense = 50000
	flag       = 5000
)

type TransactionService struct {
	tpb.UnimplementedTransactionServiceServer
	DB    *storage.Queries
	Kafka *kafka.Kafka
}

func NewTransactionService(q *storage.Queries, k *kafka.Kafka) *TransactionService {
	return &TransactionService{DB: q, Kafka: k}
}

func (t *TransactionService) CreateTransaction(ctx context.Context, req *tpb.TransactionInfo) (*tpb.TransactionID, error) {
	user_id := ctx.Value("user_id").(string)
	id := uuid.New().String()
	date := time.Now().Format(time.ANSIC)

	totalIncome, err := t.DB.GetIncomes(ctx, storage.GetIncomesParams{
		Type:   "income",
		UserID: user_id,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get total income: %v", err)
	}
	totalExpense, err := t.DB.GetExpenses(ctx, storage.GetExpensesParams{
		Type:   "expense",
		UserID: user_id,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get total expense: %v", err)
	}
	switch req.Type {
	case "income":
		if totalIncome+int64(req.Amount) >= int64(MaxIncome-flag) {
			message := "you are about to reach maximum income limit !"
			err := t.Kafka.ExceedingIncomeLimit(message, user_id, &tpb.ReportResponse{
				Income:     float64(totalIncome) + req.Amount,
				Expenses:   float64(totalExpense),
				NetSavings: float64(totalIncome) - float64(totalExpense),
			})
			if err != nil {
				return nil, fmt.Errorf("failed to send message: %v", err)
			}
		}
	case "expense":
		message := "you are about to reach maximum expense limit !"
		if totalExpense+int64(req.Amount) >= int64(MaxExpense-flag) {
			err := t.Kafka.ExceedingIncomeLimit(message, user_id, &tpb.ReportResponse{
				Income:     float64(totalIncome) + req.Amount,
				Expenses:   float64(totalExpense),
				NetSavings: float64(totalIncome) - float64(totalExpense),
			})
			if err != nil {
				return nil, fmt.Errorf("failed to send message: %v", err)
			}
		}
	}

	response, err := t.DB.CreateTransaction(ctx, storage.CreateTransactionParams{
		UserID:   user_id,
		ID:       id,
		Type:     req.Type,
		Category: req.Category,
		Currency: req.Currency,
		Amount:   req.Amount,
		Date:     date,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create new transactions: %v", err)
	}
	return &tpb.TransactionID{Id: response}, nil
}

func (t *TransactionService) UpdateTransactionByID(ctx context.Context, req *tpb.TransactionWithID) (*tpb.TransactionResponse, error) {
	date := time.Now().Format(time.ANSIC)
	err := t.DB.UpdateTransactionByID(ctx, storage.UpdateTransactionByIDParams{
		ID:       req.Id,
		Type:     req.Type,
		Category: req.Category,
		Currency: req.Currency,
		Amount:   req.Amount,
		Date:     date,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update transaction: %v", err)
	}

	return &tpb.TransactionResponse{
		Message: "Transaction updated successfully",
	}, nil
}

func (t *TransactionService) DeleteTransactionByID(ctx context.Context, req *tpb.TransactionID) (*tpb.TransactionResponse, error) {
	err := t.DB.DeleteTransactionByID(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete transaction: %v", err)
	}

	return &tpb.TransactionResponse{
		Message: "Transaction deleted successfully",
	}, nil
}

func (t *TransactionService) GetTransactionByID(ctx context.Context, req *tpb.TransactionID) (*tpb.TransactionWithID, error) {
	res, err := t.DB.GetTransactionByID(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction by id: %v", err)
	}

	return &tpb.TransactionWithID{
		Id:       res.ID,
		Type:     res.Type,
		Category: res.Category,
		Currency: res.Currency,
		Amount:   res.Amount,
		Date:     res.Date,
	}, nil
}

func (t *TransactionService) GetTransactionsByCategory(ctx context.Context, req *tpb.TransactionCategory) (*tpb.ListTransactions, error) {
	user_id := ctx.Value("user_id").(string)
	res, err := t.DB.GetTransactionsByCategory(ctx, storage.GetTransactionsByCategoryParams{
		UserID:   user_id,
		Category: req.Category,
	})
	if err != nil {
		fmt.Println("exxx xato")
		return nil, fmt.Errorf("failed to get transaction by category: %v", err)
	}

	var transactions []*tpb.TransactionWithID
	for _, t := range res {
		transactions = append(transactions, &tpb.TransactionWithID{
			Id:       t.ID,
			Type:     t.Type,
			Category: t.Category,
			Currency: t.Currency,
			Amount:   t.Amount,
			Date:     t.Date,
		})
	}

	return &tpb.ListTransactions{
		ListTransactions: transactions,
	}, nil
}

func (t *TransactionService) GetTransactionByDate(ctx context.Context, req *tpb.TransactionDate) (*tpb.ListTransactions, error) {
	user_id := ctx.Value("user_id").(string)
	res, err := t.DB.GetTransactionByDate(ctx, storage.GetTransactionByDateParams{
		UserID: user_id,
		Type:   req.Type,
		Date:   req.Start,
		Date_2: req.End,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction by date: %v", err)
	}
	var transactions []*tpb.TransactionWithID
	for _, t := range res {
		transactions = append(transactions, &tpb.TransactionWithID{
			Id:       t.ID,
			Type:     t.Type,
			Category: t.Category,
			Currency: t.Currency,
			Amount:   t.Amount,
			Date:     t.Date,
		})
	}
	return &tpb.ListTransactions{
		ListTransactions: transactions,
	}, nil
}

func (t *TransactionService) GetReport(ctx context.Context, req *tpb.Empty) (*tpb.ReportResponse, error) {
	user_id := ctx.Value("user_id").(string)
	total_income, err := t.DB.GetIncomes(ctx, storage.GetIncomesParams{
		UserID: user_id,
		Type:   "income",
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get reports: %v", err)
	}

	total_expense, err := t.DB.GetExpenses(ctx, storage.GetExpensesParams{
		UserID: user_id,
		Type:   "expense",
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get reports: %v", err)
	}

	return &tpb.ReportResponse{
		Income:     float64(total_income),
		Expenses:   float64(total_expense),
		NetSavings: float64(total_income) - float64(total_expense),
	}, nil
}
