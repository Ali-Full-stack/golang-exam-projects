package router

import (
	"gateway/client"
	"gateway/api/handler"
	"gateway/storage/redisdb"
	"os"


	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_"gateway/api/swagger/docs"
	swaggerFiles "github.com/swaggo/files"
	swag "github.com/swaggo/gin-swagger"
)

// New ...
// @title  PERSONAL FINANCE MANAGEMENT
// @description This swagger UI was created to manage personal finance
// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:9000
// @contact.email  ali.team@gmail.com

func Routes() {
	r := gin.Default()

	userClient := client.DialUsersClient(os.Getenv("users_url"))
	budgetClient := client.DialBudgetClient(os.Getenv("budget_url"))
	redis := redisdb.ConnectRedis(os.Getenv("redis_url"))
	incomeClient := client.DialIncomeClient(os.Getenv("income_url"))
	notificationClient := client.DialNotifyClient("notification_url")
	Handler := handler.NewHandler(userClient, redis, budgetClient, incomeClient, notificationClient)

	// //USERS
	// r.POST("/api/users/register", Handler.RegisterUser)
	// r.POST("/api/users/verification", Handler.VerifyCode)
	// r.POST("/api/users/login", Handler.UserLogin)
	// r.POST("/api/users/forgot-password", Handler.ForgetPassword)
	// r.PUT("/api/users/update-password", Handler.UpdatePassword)
	// r.DELETE("/api/users/logout", Handler.UserLogout)
	// r.GET("/api/users/:id", Handler.GetUserByID)

	//BUDGETS
	r.POST("/api/budgets/create", Handler.CreateBudget)
	r.PUT("/api/budgets", Handler.UpdateBudgetAmount)
	r.GET("/api/budgets", Handler.GetBudgets)
	r.DELETE("/api/budgets/:id", Handler.DeleteBudgetByID)

	// //Income
	// r.POST("/api/income", Handler.CreateTransaction)
	// r.PUT("api/income/:id", Handler.UpdateTransactionByID)
	// r.DELETE("api/incomde/:id", Handler.DeleteTransactionByID)
	// r.GET("api/income/:id", Handler.GetTransactionByID)
	// r.GET("api/income/category/:category", Handler.GetTransactionsByCategory)
	// r.GET("api/income/date", Handler.GetTransactionByDate)

	// //Notify
	// r.GET("api/notifications", Handler.GetNotification)
	// r.GET("api/notifications/unread", Handler.GetUnreadNotifications)

	r.GET("/swagger/*any", swag.WrapHandler(swaggerFiles.Handler))
	
	r.Run(os.Getenv("gateway_url"))

}
