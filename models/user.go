package models

import "time"

// USER MANAGEMENT : Registration, login, roles (admin/user)

type User struct {
	ID         string
	Name       string
	Email      string
	Username   string
	Password   string
	Address    Residence
	Created_at time.Time
}

type Residence struct {
	ZipCode     string
	City        string
	State       string
	Country     string
	Street      string
	HouseNumber string
}
