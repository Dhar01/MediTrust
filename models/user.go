package models

import "github.com/google/uuid"

type User struct {
	Address      Address
	Age          int32
	Email        string
	HashPassword string
	Id           uuid.UUID
	IsActive     bool
	Name         FullName
	Phone        string
	Role         string
}

type FullName struct {
	FirstName string
	LastName  string
}

type Address struct {
	City          string
	Country       string
	PostalCode    string
	StreetAddress string
}

type FetchUserInfoResponse struct {
	Address  Address
	Age      int32
	Email    string
	IsActive bool
	Name     FullName
	Phone    string
	Role     string
}

type UpdateUserRequest struct {
	Address Address
	Age     int32
	Name    FullName
	Phone   string
	Email   string
}

type UpdateUserResponse struct {
	Address  Address
	Age      int32
	Email    string
	IsActive bool
	Name     FullName
	Phone    string
	Role     string
}

type LogInResp struct {
	AccessToken  string
	RefreshToken string
}
