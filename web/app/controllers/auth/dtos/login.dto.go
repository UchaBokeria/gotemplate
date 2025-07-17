package dtos

type LoginDto struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6"`
}

type RegisterDto struct {
	Name     string `form:"name" validate:"required,min=2"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6"`
}
