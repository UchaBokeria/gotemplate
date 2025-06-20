package schemas

type Login struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type Register struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
