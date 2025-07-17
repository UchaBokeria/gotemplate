package translations

import (
	"main/internal/models"
	"main/internal/storage"
)

var Seed = []models.Translation{
	// English translations
	{Key: "welcome", Value: "Welcome to Alfashop", Language: "en"},
	{Key: "products", Value: "Products", Language: "en"},
	{Key: "about", Value: "About Us", Language: "en"},
	{Key: "contact", Value: "Contact", Language: "en"},
	{Key: "home", Value: "Home", Language: "en"},
	{Key: "categories", Value: "Categories", Language: "en"},
	{Key: "cart", Value: "Shopping Cart", Language: "en"},
	{Key: "search", Value: "Search products...", Language: "en"},
	{Key: "price", Value: "Price", Language: "en"},
	{Key: "add_to_cart", Value: "Add to Cart", Language: "en"},
	{Key: "checkout", Value: "Checkout", Language: "en"},
	{Key: "login", Value: "Login", Language: "en"},
	{Key: "register", Value: "Register", Language: "en"},
	{Key: "air_filters", Value: "Air Filters", Language: "en"},
	{Key: "oil_filters", Value: "Oil Filters", Language: "en"},
	{Key: "fuel_filters", Value: "Fuel Filters", Language: "en"},
	{Key: "brake_parts", Value: "Brake Parts", Language: "en"},

	// Russian translations
	{Key: "welcome", Value: "Добро пожаловать в Alfashop", Language: "ru"},
	{Key: "products", Value: "Продукты", Language: "ru"},
	{Key: "about", Value: "О нас", Language: "ru"},
	{Key: "contact", Value: "Контакты", Language: "ru"},
	{Key: "home", Value: "Главная", Language: "ru"},
	{Key: "categories", Value: "Категории", Language: "ru"},
	{Key: "cart", Value: "Корзина", Language: "ru"},
	{Key: "search", Value: "Поиск товаров...", Language: "ru"},
	{Key: "price", Value: "Цена", Language: "ru"},
	{Key: "add_to_cart", Value: "В корзину", Language: "ru"},
	{Key: "checkout", Value: "Оформить заказ", Language: "ru"},
	{Key: "login", Value: "Войти", Language: "ru"},
	{Key: "register", Value: "Регистрация", Language: "ru"},
	{Key: "air_filters", Value: "Воздушные фильтры", Language: "ru"},
	{Key: "oil_filters", Value: "Масляные фильтры", Language: "ru"},
	{Key: "fuel_filters", Value: "Топливные фильтры", Language: "ru"},
	{Key: "brake_parts", Value: "Тормозные детали", Language: "ru"},

	// Georgian translations
	{Key: "welcome", Value: "კეთილი იყოს თქვენი მობრძანება Alfashop-ში", Language: "ge"},
	{Key: "products", Value: "პროდუქტები", Language: "ge"},
	{Key: "about", Value: "ჩვენს შესახებ", Language: "ge"},
	{Key: "contact", Value: "კონტაქტი", Language: "ge"},
	{Key: "home", Value: "მთავარი", Language: "ge"},
	{Key: "categories", Value: "კატეგორიები", Language: "ge"},
	{Key: "cart", Value: "კალათა", Language: "ge"},
	{Key: "search", Value: "მოძებნეთ პროდუქტები...", Language: "ge"},
	{Key: "price", Value: "ფასი", Language: "ge"},
	{Key: "add_to_cart", Value: "კალათაში დამატება", Language: "ge"},
	{Key: "checkout", Value: "შეკვეთა", Language: "ge"},
	{Key: "login", Value: "შესვლა", Language: "ge"},
	{Key: "register", Value: "რეგისტრაცია", Language: "ge"},
	{Key: "air_filters", Value: "ჰაერის ფილტრები", Language: "ge"},
	{Key: "oil_filters", Value: "ზეთის ფილტრები", Language: "ge"},
	{Key: "fuel_filters", Value: "საწვავის ფილტრები", Language: "ge"},
	{Key: "brake_parts", Value: "სამუხრუჭე დეტალები", Language: "ge"},
}

func Populate() {
	for _, row := range Seed {
		storage.DB.Create(&row)
	}
}
