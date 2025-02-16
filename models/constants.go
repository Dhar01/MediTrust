package models

var (
	Admin    string = "admin"
	Customer string = "customer"

	Email string = "email"
	Phone string = "phone"

	RootPath     string = "/"
	TokenRefresh string = "refresh_token"
	TokenNull    string = ""

	Dev string = "dev"

	/*
		TODO: A cmd app to insert the value below or manually edit this file or move it to config
	*/
	DomainName  string = "localhost"
	CompanyName string = "company"
	DomainPort  int    = 8080
)
