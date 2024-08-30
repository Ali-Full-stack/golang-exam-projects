package handler

import (
	"context"
	"encoding/json"
	"gateway/internal/model"
	upb "gateway/protos/user"
	"io"
	"log"
	"net/http"
)

type UserHandler struct {
	Client upb.UserServiceClient
}

func NewUserHandler(u upb.UserServiceClient) *UserHandler {
	return &UserHandler{Client: u}
}
// @Router  				/api/users/register [post]
// @Summary 			Registers New User
// @Description 		This method used  to register new  users 
// @Security 				BearerAuth
// @Tags					 USERS
// @accept					json
// @Produce				  json
// @Param 					body    body    model.UserInfo    true  "User Information"
// @Success					201 	{object}   user.UserID		"User ID"
// @Failure					 400 {object} error "Invalid Request Body"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (u *UserHandler) RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	var userJS model.UserInfo
	if err := json.NewDecoder(r.Body).Decode(&userJS); err != nil {
		log.Println("failed to decode request body:", err)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	userPB := &upb.UserInfo{
		Username: userJS.Username,
		Email:    userJS.Email,
		Password: userJS.Password,
	}

	resp, err := u.Client.RegisterUser(r.Context(), userPB)
	if err != nil {
		log.Println("failed to get grpc response on Register User :", err)
		http.Error(w, "Unable to get response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
// @Router  				/api/users/login [post]
// @Summary 			User Login
// @Description 		This method used  for users to login  
// @Security 				BearerAuth
// @Tags					 USERS
// @accept					json
// @Produce				  json
// @Param 					body    body    model.UserLogin    true  "User login"
// @Success					201 	{object}   user.UserToken		"User Token"
// @Failure					 400 {object} error "Invalid Request Body"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (u *UserHandler) UserLogin(w http.ResponseWriter, r *http.Request) {
	var userJS model.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&userJS); err != nil {
		log.Println("failed to decode request body:", err)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	userPB := &upb.UserLogin{
		Id:       userJS.Id,
		Password: userJS.Password,
	}
	token, err := u.Client.LoginUser(r.Context(), userPB)
	if err != nil {
		log.Println("failed to get grpc response on Login User :", err)
		http.Error(w, "Unable to get response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
// @Router  				/api/users/{id} [PUT]
// @Summary 			Update User Information
// @Description 		This method used  to update users information  
// @Security 				BearerAuth
// @Tags					 USERS
// @accept					json
// @Produce				  json
// @Param 					id    	path        string    true    "User  ID"
// @Param 					body    body    model.UserInfo    true  "User Information"
// @Success					201 	{object}   user.UserResponse		"Response"
// @Failure					 400 {object} error "Invalid Request Body"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ctx := context.WithValue(r.Context(), "id", id)
	var userJS model.UserInfo
	if err := json.NewDecoder(r.Body).Decode(&userJS); err != nil {
		log.Println("failed to decode request body:", err)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	userPB := &upb.UserInfo{
		Username: userJS.Username,
		Email:    userJS.Email,
		Password: userJS.Password,
	}
	resp, err := u.Client.UpdateUser(ctx, userPB)
	if err != nil {
		log.Println("failed to get grpc response on Update User :", err)
		http.Error(w, "Unable to get response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
// @Router  				/api/users/{id} [DELETE]
// @Summary 			Delete User Information
// @Description 		This method used  to Delete users information  
// @Security 				BearerAuth
// @Tags					 USERS
// @accept					json
// @Produce				  json
// @Param 					id    	path        string    true    "User  ID"
// @Success					201 	{object}   user.UserResponse		"Response"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (u *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	resp, err := u.Client.DeleteUser(r.Context(), &upb.UserID{Id: id})
	if err != nil {
		log.Println("failed to get grpc response on Delete User :", err)
		http.Error(w, "Unable to get response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
// @Router  				/api/users [get]
// @Summary 			Get All Users
// @Description 		This method used  to get all users   
// @Security 				BearerAuth
// @Tags					 USERS
// @accept					json
// @Produce				  json
// @Success					201 	{object}   []user.UserWithID		"User Information"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (u *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	stream, err := u.Client.GetAllUsers(r.Context(), &upb.Empty{})
	if err != nil {
		log.Println("failed to get grpc response on Get All User :", err)
		http.Error(w, "Unable to get response", http.StatusInternalServerError)
		return
	}
	var users []*upb.UserWithID
	for {
		user, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("failed streaming  on Get All User :", err)
			http.Error(w, "Unable to get response", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}
// @Router  				/api/users/{id} [get]
// @Summary 			Get User By ID
// @Description 		This method used  to Get User By ID  
// @Security 				BearerAuth
// @Tags					 USERS
// @accept					json
// @Produce				  json
// @Param 					id    	path        string    true    "Client ID"
// @Success					201 	{object}   user.UserWithID		"User ID"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (u *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	resp, err := u.Client.GetUserById(r.Context(), &upb.UserID{Id: id})
	if err != nil {
		log.Println("failed to get grpc response on Get User By ID :", err)
		http.Error(w, "Unable to get response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
