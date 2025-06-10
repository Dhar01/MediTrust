package model

import "github.com/google/uuid"

// User defines the structure of user profile
type User struct {
	ID      uint64
	IDAuth  uuid.UUID
	Name    FullName
	Address Address
	Age     int32
}

// FullName defines the structure of users name
type FullName struct {
	FirstName string
	LastName  string
}

// Address defines the user's address
type Address struct {
	Country string
	City    string
	RoadNo  string
	HouseNo string
	FlatNo  string
}
