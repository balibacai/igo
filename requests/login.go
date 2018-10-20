package requests

type LoginCredentials struct {
	Email    string `form:"email"valid:"Required;Email;MaxSize(32)"`
	Password string `form:"password"valid:"Required;MaxSize(64)"`
}
