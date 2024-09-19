package handler

import (
	"context"
	"database/sql"
	"fmt"
	"user-service/internal/auth/hash"
	"user-service/internal/auth/token"
	upb "user-service/proto"
	"user-service/storage"

	"github.com/google/uuid"
)

type UserService struct {
	upb.UnimplementedUserServiceServer
	DB *storage.Queries
}

func NewUserService(q *storage.Queries) *UserService {
	return &UserService{DB: q}
}

func (u *UserService) CreateUser(ctx context.Context, req *upb.UserInfo) (*upb.UserWithID, error) {
	id := uuid.New().String()
	hash_password, err := hash.GenerateHash(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to generate hash from userpassword: %v", err)
	}

	result, err := u.DB.CreateUser(ctx, storage.CreateUserParams{
		ID:   id,
		Name: req.Name,
		Email: sql.NullString{
			Valid:  true,
			String: req.Email,
		},
		PasswordHash: hash_password,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to add user into postgres: %v", err)
	}
	return &upb.UserWithID{Id: result.ID, Name: result.Name, Email: result.Email.String}, nil
}

func (u *UserService) LoginUser(ctx context.Context, req *upb.UserLogin) (*upb.UserToken, error) {
	result, err := u.DB.LoginUser(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("unable to login user: %v", err)
	}
	token, err := token.GenerateToken(req.Id, result)
	if err != nil {
		return nil, fmt.Errorf("unable to generate token: %v", err)
	}
	return &upb.UserToken{Token: token}, nil
}

func (u *UserService) ForgotPassword(ctx context.Context, req *upb.UserID) (*upb.UserEmail, error) {
	result, err := u.DB.ForgotPassword(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("unable to get user email: %v", err)
	}
	return &upb.UserEmail{Email: result.String}, nil
}

func (u *UserService) UpdatePassword(ctx context.Context, req *upb.UserPassword) (*upb.UserResponse, error) {
	hash_password, err := hash.GenerateHash(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to generate hash from userpassword: %v", err)
	}
	err = u.DB.UpdatePassword(ctx, storage.UpdatePasswordParams{
		Email:        sql.NullString{String: req.GetEmail(), Valid: true},
		PasswordHash: hash_password,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to update password: %v", err)
	}
	return &upb.UserResponse{Message: "Password updated successfully"}, nil
}

func (u *UserService) LogOutUser(ctx context.Context, req *upb.UserID) (*upb.UserResponse, error) {
	err := u.DB.LogOutUser(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("unable to logout user: %v", err)
	}
	return &upb.UserResponse{Message: "user logged out successfully"}, nil
}

func (u *UserService) GetUserById(ctx context.Context, req *upb.UserID) (*upb.UserWithID, error) {
	result, err := u.DB.GetUserByID(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("unable to get user by id: %v", err)
	}
	return &upb.UserWithID{Id: result.ID, Name: result.Name, Email: result.Email.String}, nil
}
