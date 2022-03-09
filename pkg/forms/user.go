package forms

// LoginForm is used to log the user in
type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
