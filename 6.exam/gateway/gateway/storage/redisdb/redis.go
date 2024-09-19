package redisdb

import (
	"context"
	"fmt"
	"gateway/internal/model"
	"gateway/proto/users"
	"strconv"
	"time"

	r "github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *r.Client
}

func ConnectRedis(redis_url string) *RedisClient {
	return &RedisClient{Client: r.NewClient(&r.Options{Addr: redis_url})}
}

func (r *RedisClient) AddUserIntoRedis(ctx context.Context, req *users.UserInfo, code int64) error {
	err := r.Client.HSet(ctx, req.Email, map[string]interface{}{
		"name":     req.Name,
		"password": req.Password,
		"code":     code,
	}).Err()
	if err != nil {
		return fmt.Errorf("error HSET:  %v", err)
	}
	r.Client.Expire(ctx, req.Email, 5*time.Minute)
	return nil
}

func (r *RedisClient) VerifyCodeAndGetUser(ctx context.Context, req model.UserCode) (*users.UserInfo, error) {

	result, err := r.Client.HGetAll(ctx, req.Email).Result()
	if err != nil {
		return nil, fmt.Errorf("unable to get user info from redis :%v", err)
	}
	code, _ := strconv.Atoi(result["code"])
	if code != int(req.Code) {
		return nil, fmt.Errorf("code is Invalid ")
	}
	return &users.UserInfo{
		Name:     result["name"],
		Email:    result["email"],
		Password: result["password"],
	}, nil

}

func (r *RedisClient) ForgetPasswordIntoRedis(ctx context.Context, email string, code int64) error {
	data := map[string]interface{}{
		"code": fmt.Sprintf("%d", code),
	}

	err := r.Client.HSet(ctx, email, data).Err()
	if err != nil {
		return fmt.Errorf("error HSET: %v", err)
	}
	r.Client.Expire(ctx, email, 5*time.Minute)
	return nil
}

func (r *RedisClient) VerifyForgetPasswordUser(ctx context.Context, req model.NewPass) (*users.UserPassword, error) {
	result, err := r.Client.HGetAll(ctx, req.Email).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get user info while Redis: %v", err)
	}
	codeStr, exists := result["code"]
	if !exists {
		return nil, fmt.Errorf("code is Invalid ")
	}

	code, err := strconv.Atoi(codeStr)
	if err != nil {
		return nil, fmt.Errorf("failed to convert code to int: %v", err)
	}
	if code != int(req.Code) {
		return nil, fmt.Errorf("invalid code")
	}

	return &users.UserPassword{
		Email:    req.Email,
		Password: result["password"],
	}, nil
}

// func (r *RedisClient) CheckUserForLogin(ctx context.Context, user *upb.UserLogin) (*upb.UserToken, error) {

// 	result, err := r.Client.HGetAll(ctx, user.Id).Result()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get user info while LOGIN: %v", err)
// 	}
// 	if result["password"] != user.Password {
// 		return nil, fmt.Errorf("invalid User Password : %v", err)
// 	}
// 	token, err := auth.GenerateToken(user.Id, "user")
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to generate user Token")
// 	}
// 	return token, nil
// }

// func (r *RedisClient) UpdateUserInRedis(id string, user *upb.UserInfo) error {
// 	err := r.Client.HSet(context.Background(), id, map[string]interface{}{
// 		"password": user.Password,
// 	}).Err()
// 	if err != nil {
// 		return fmt.Errorf("error HSET:  %v", err)
// 	}
// 	return nil
// }

// func (r *RedisClient) DeleteUserFromRedis(user *upb.UserID) error {
// 	_, err := r.Client.Del(context.Background(), user.Id, user.Id).Result()
// 	if err != nil {
// 		return fmt.Errorf("error DEL:%v", err)
// 	}
// 	return nil
// }
