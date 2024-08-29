package postgres

import (
	"fmt"
	"user-service/auth"
	"user-service/internal/model"
	upb "user-service/protos"

	sqrl "github.com/Masterminds/squirrel"
)

func (p *Postgres) AddUserIntoPostgres(id string, user *upb.UserInfo) error {
	hashPassword, err := auth.GenerateHashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("in postgres : %v", err)
	}
	query, args, err := sqrl.
		Insert("users").
		Columns("id, username, email, password").
		Values(id, user.Username, user.Email, hashPassword).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("failed query in INSERT: %v", err)
	}

	_, err = p.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to exec user information: %v", err)
	}
	return nil
}

func (p *Postgres) UpdateUserInPostgres(id string, user *upb.UserInfo) (*upb.UserResponse, error) {
	password, err := auth.GenerateHashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("faield to generate hashpassword on UPDATE: %v", err)
	}
	query, args, err := sqrl.Update("users").
		SetMap(map[string]interface{}{
			"username": user.Username,
			"email":    user.Email,
			"password": password,
		}).Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed query on UPDATE: %v", err)
	}
	result, err := p.DB.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update user in postgres:%v", err)
	}
	if n, _ := result.RowsAffected(); n == 0 {
		return nil, fmt.Errorf("user does not exists with ID %v", id)
	}
	return &upb.UserResponse{Message: "User Info Updated succesfully"}, nil
}

func (p *Postgres) DeleteUserFromPostgres(user *upb.UserID) (*upb.UserResponse, error) {
	query, args, err := sqrl.Delete("*").
		From("users").
		Where(sqrl.Eq{"id": user.Id}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed query on DELETE: %v", err)
	}
	result, err := p.DB.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to delete user in postgres:%v", err)
	}
	if n, _ := result.RowsAffected(); n == 0 {
		return nil, fmt.Errorf("user does not exists with ID %v", user.Id)
	}
	return &upb.UserResponse{Message: "User Info Deleted succesfully"}, nil
}

func (p *Postgres) GetAllUsersFromPostgres(req *upb.Empty)([]*upb.UserWithID, error){
	query, args, err :=sqrl.Select("id, username, email").
	From("users").
	PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed query on Select All Users: %v",err)
	}
	rows, err :=p.DB.Query(query,args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %v",err)
	} 
	defer rows.Close()
	var users []*upb.UserWithID

	for rows.Next(){
		var u model.UserWithID
		if err :=rows.Scan(&u.Id, &u.Username, &u.Email); err != nil {
			return nil, fmt.Errorf("failed to scan on Getting All Users : %v",err)
		}
		users = append(users, &upb.UserWithID{
			Id: u.Id,
			Username: u.Username,
			Email: u.Email,
		})
	}
	return users, nil
}

func (p *Postgres) GetUserByIdFromPostgres(req *upb.UserID)(*upb.UserWithID, error){
	query, args, err :=sqrl.Select("id, username, email"). 
	From("users"). 
	Where(sqrl.Eq{"id" : req.Id}). 
	PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed query on getting users by id : %v",err)
	}
	row :=p.DB.QueryRow(query, args...)

	var user model.UserWithID
	if err :=row.Scan(&user.Id, &user.Username, &user.Email); err != nil {
		return nil, fmt.Errorf("failed to scan on Getting User By ID: %v",err)
	}
	return &upb.UserWithID{Id: user.Id, Username: user.Username, Email: user.Email}, nil
}