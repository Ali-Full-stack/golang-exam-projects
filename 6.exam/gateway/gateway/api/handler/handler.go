package handler

import (
	"gateway/proto/budget"
	"gateway/proto/income"
	"gateway/proto/notification"
	"gateway/proto/users"
	"gateway/storage/redisdb"
)

type Handler struct {
	User  users.UserServiceClient
	Redis *redisdb.RedisClient
	Budget budget.BudgetServiceClient
	Income income.TransactionServiceClient
	Notification notification.NotificationServiceClient
}

func NewHandler(u users.UserServiceClient, r *redisdb.RedisClient, b budget.BudgetServiceClient, i income.TransactionServiceClient, n notification.NotificationServiceClient)*Handler{
	return &Handler{
		User: u,
		Redis: r,
		Budget: b,
		Income: i,
		Notification: n,
	}
}