package models

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type AdminUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Role      string `json:"role"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type AdminPost struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	ImageUrl string `json:"image-url"`
	Category string `json:"category"`
}

type Post struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	ImageUrl string `json:"image-url"`
	Category string `json:"category"`
}

type Comment struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	PostId      string `json:"postId"`
}
