package redisdb

import (
	"context"
	"fmt"
	"user-service/auth"
	upb "user-service/protos"

	r "github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *r.Client
}

func ConnectRedis(redis_url string) *RedisClient {
	return &RedisClient{Client: r.NewClient(&r.Options{Addr: redis_url})}
}

func (r *RedisClient) AddUserIntoRedis(ctx context.Context, id, password string) error {
	err := r.Client.HSet(ctx, id, map[string]interface{}{
		"password": password,
	}).Err()
	if err != nil {
		return fmt.Errorf("error HSET:  %v", err)
	}
	return nil
}

func (r *RedisClient) CheckUserForLogin(ctx context.Context, user *upb.UserLogin) (*upb.UserToken, error) {

	result, err := r.Client.HGetAll(ctx, user.Id).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get user info while LOGIN: %v", err)
	}
	if result["password"] != user.Password {
		return nil, fmt.Errorf("invalid User Password : %v", err)
	}
	token, err := auth.GenerateToken(user.Id, "user")
	if err != nil {
		return nil, fmt.Errorf("failed to generate user Token")
	}
	return token, nil
}

func (r *RedisClient) UpdateUserInRedis(id string, user *upb.UserInfo)error {
	err := r.Client.HSet(context.Background(), id, map[string]interface{}{
		"password": user.Password,
	}).Err()
	if err != nil {
		return fmt.Errorf("error HSET:  %v", err)
	}
	return nil
}

func (r *RedisClient) DeleteUserFromRedis(user *upb.UserID)error{
	_,err :=r.Client.Del(context.Background(), user.Id, user.Id).Result()
	if err != nil {
		return fmt.Errorf("error DEL:%v",err)
	}
	return nil
}