package models

type User struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Picture   string `json:"picture,omitempty"`
	Role      string `json:"role,omitempty"`
	Password  string `json:"password,omitempty"`
}

type UserRegister struct {
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	Email   string `json:"email,omitempty"`
	// Picture  string `json:"picture,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserLogin struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
