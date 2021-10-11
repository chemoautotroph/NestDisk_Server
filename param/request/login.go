package request

type Login struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
