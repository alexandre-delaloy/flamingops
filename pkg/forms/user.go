package forms

// LoginForm is used to log the user in
type LoginForm struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}
