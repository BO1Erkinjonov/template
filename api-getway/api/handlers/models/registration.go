package models

type RegisterModel struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type RegisterResponse struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	UserName     string `json:"username"`
	Role         string `json:"role"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Id           string `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
