package auth

import (
	"main/internal/models"
	"main/internal/storage"
	"main/web/app/types"
	"main/web/app/view/components"

	"github.com/UchaBokeria/goyard/controller"
	"golang.org/x/crypto/bcrypt"
)

type LoginDto struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type RegisterDto struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func Login(ctx *controller.Context, loginData *LoginDto) error {
	// Validate input
	if loginData.Email == "" || loginData.Password == "" {
		return ctx.Html(components.ProductNotFound(ctx.Get("webconfig").(types.WebConfig)))
	}

	// Find user by email
	var user models.Users
	if err := storage.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		return ctx.Html(components.ProductNotFound(ctx.Get("webconfig").(types.WebConfig)))
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return ctx.Html(components.ProductNotFound(ctx.Get("webconfig").(types.WebConfig)))
	}

	// Set session
	ctx.Set("user_id", user.ID)
	ctx.Set("user_email", user.Email)

	return ctx.Html(components.ProductNotFound(ctx.Get("webconfig").(types.WebConfig)))
}

func Register(ctx *controller.Context, registerData *RegisterDto) error {
	// Validate input
	if registerData.Name == "" || registerData.Email == "" || registerData.Password == "" {
		return ctx.Html(components.ProductNotFound(ctx.Get("webconfig").(types.WebConfig)))
	}

	// Check if user already exists
	var existingUser models.Users
	if err := storage.DB.Where("email = ?", registerData.Email).First(&existingUser).Error; err == nil {
		return ctx.Html(components.ProductNotFound(ctx.Get("webconfig").(types.WebConfig)))
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.Html(components.ProductNotFound(ctx.Get("webconfig").(types.WebConfig)))
	}

	// Create user
	user := models.Users{
		Fullname: registerData.Name,
		Email:    registerData.Email,
		Password: string(hashedPassword),
	}

	if err := storage.DB.Create(&user).Error; err != nil {
		return ctx.Html(components.ProductNotFound(ctx.Get("webconfig").(types.WebConfig)))
	}

	// Set session
	ctx.Set("user_id", user.ID)
	ctx.Set("user_email", user.Email)

	return ctx.Html(components.ProductNotFound(ctx.Get("webconfig").(types.WebConfig)))
}

func Logout(ctx *controller.Context) error {
	// Clear session
	ctx.Set("user_id", nil)
	ctx.Set("user_email", nil)

	return ctx.Html(components.ProductNotFound(ctx.Get("webconfig").(types.WebConfig)))
}
