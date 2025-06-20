package interfaces

import (
	"main/internal/models"
	"main/internal/storage"
)

func Pint(num int) *int {
	return &num
}

var Slideshows = []models.Interface_slideShow{
	{
		InterfaceID: 1,
		Name:        "Slideshow 2",
		Slug:        "slideshow-2",
		Slogan:      "alfashop ხდება MERCEDES-AMG MOTORSPORT- ის „ოფიციალური მიმწოდებელი“",
		Desc:        `აფფალტერბახი/ვერტერი. 100% დამზადებულია გერმანიაში - Mercedes-AMG და alfashop ერთად წავლენ იპოდრომზე მომავალში: Ravensberger Gleitstoffvertrieb GmbH-ის სასაქონლო ნიშანი ახლა არის "ოფიციალური მიმწოდებელი" მომხმარებელთა სარბოლო პროგრამაში Affalterbach-ის პერფორმანსისა და სპორტული ავტომობილების ბრენდის`,
		Url:         "/posts/2",
		TypeID:      1,
		PicID:       Pint(2),
		Index:       2,
	},
	{
		InterfaceID: 1,
		Name:        "Slideshow 1",
		Slug:        "slideshow-1",
		Slogan:      "მანქანა ფორმულა 1-ის 2024 წლის სეზონისთვის გამოვლინდა!",
		Desc:        "ახალი ლაივერი მსოფლიო საზოგადოებას ლას-ვეგასში გაშვების ღონისძიებაზე წარუდგინა - თანამედროვე დიზაინში დომინირებს ლურჯი ჩრდილები. ეს მომხიბლავი მანქანა იქნება მძღოლების იუკი ცუნოდასა და დანიელ რიკიარდოს ახალი სამუშაო ადგილი, რომლებიც მოუთმენლად ელიან სეზონის დაწყებას 29 თებერვალს ბაჰრეინის გრან პრიზე.",
		Url:         "/posts/1",
		TypeID:      1,
		PicID:       Pint(1),
		Index:       1,
	},
	{
		InterfaceID: 1,
		Name:        "Slideshow 1",
		Slug:        "slideshow-1",
		Slogan:      "",
		Desc:        "",
		Url:         "#",
		TypeID:      1,
		PicID:       Pint(3),
		Index:       3,
	},
}

func Slideshow() {
	for _, row := range Slideshows {
		storage.DB.Create(&row)
	}
}
