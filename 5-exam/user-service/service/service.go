package service

import (
	"context"
	"fmt"
	"user-service/internal/postgres"
	"user-service/internal/redisdb"
	"user-service/kafka"
	upb "user-service/protos"

	"github.com/google/uuid"
)

type UserService struct {
	upb.UnimplementedUserServiceServer
	RedisClient *redisdb.RedisClient
	Postgres    *postgres.Postgres
	Kafka       *kafka.Kafka
}

func NewUserService(r *redisdb.RedisClient, p *postgres.Postgres, k *kafka.Kafka) *UserService {
	return &UserService{RedisClient: r, Postgres: p, Kafka: k}
}

func (u *UserService) RegisterUser(ctx context.Context, req *upb.UserInfo) (*upb.UserID, error) {
	id := uuid.New().String()

	if err :=u.Postgres.AddUserIntoPostgres(id, req); err  != nil {
		return nil, err
	}
	if err :=u.RedisClient.AddUserIntoRedis(ctx, id, req.Password); err != nil {
		return nil, err	
	}
	if err :=u.Kafka.ProduceRegistrationEmail(id, req); err != nil {
		return nil, err
	}
	return &upb.UserID{Id: id}, nil
}

func (u *UserService) LoginUser(ctx context.Context, req *upb.UserLogin) (*upb.UserToken, error) {

	token, err := u.RedisClient.CheckUserForLogin(ctx, req)
	if err != nil {
		return nil, err
	}
	return token, nil
}
func (u *UserService) UpdateUser(ctx context.Context, req *upb.UserInfo) (*upb.UserResponse, error) {
	id :=ctx.Value("id").(string)
	response, err :=u.Postgres.UpdateUserInPostgres(id, req)
	if err != nil {
		return nil, err
	}
	err =u.RedisClient.UpdateUserInRedis(id, req)
	if err !=nil {
		return nil, err
	}
	return response ,nil
}

func (u *UserService) DeleteUser(ctx context.Context, req *upb.UserID) (*upb.UserResponse, error) {
	response, err :=u.Postgres.DeleteUserFromPostgres(req)
	if err != nil {
		return nil, err
	}
	if err :=u.RedisClient.DeleteUserFromRedis(req); err != nil {
		return nil, err
	}
	return response ,nil
}

func (u *UserService) GetAllUsers(req *upb.Empty, stream upb.UserService_GetAllUsersServer) error {
	users, err :=u.Postgres.GetAllUsersFromPostgres(req)
	if err != nil {
		return err
	}
	for _, user :=range users {
		if err :=stream.Send(user); err != nil {
			return fmt.Errorf("failed to send grpc response on Get All Users: %v",err)
		}
	}
	return nil
}
func (u *UserService) GetUserById(ctx context.Context, req *upb.UserID) (*upb.UserWithID, error) {
	user, err :=u.Postgres.GetUserByIdFromPostgres(req)
	if err !=nil {
		return nil, err
	}
	return user, nil
}
