package interfaces

import (
	"main/internal/models"
	"main/internal/storage"
)

var Interfaces = []models.Interface{
	{
		Ver:  1,
		Name: "Configuration",
		Slug: "configuration",
		Presentation: `<h2 class="text-4xl text-center font-semibold font-nino text-primary">
            კეთილი იყოს თქვენი მობრძანება
            <b class="font-bold text-5xl font-poppins">alfashop</b>-ში
        </h2>

        <p class="w-[80%] ml-[10%] text-center text-primary py-10  font-arial">როგორც ავტომობილების პრემიუმ საპოხი მასალების და სითხეების მწარმოებელი და მიმწოდებელი, ჩვენ გთავაზობთ გადაწყვეტას ყველა საჭიროებისთვის alfashop-ის პროდუქტებით</p>
    
        <div class="absolute w-[100vw] ml-[-200px] bg-[#D7E3F4] px-[200px] py-[3rem] flex flex-wrap  justify-start items-start ">
            <div class="w-[60%] flex justify-start items-start">
                <video  class="w-full"
                        src="https://www.alfashop.de/storage/app/media/video/alfashop_Corporate_Film.de.mp4"
                        muted
                        controls="controls" 
                        autoplay="autoplay"/>
            </div>

            <div class="w-[40%] h-full px-5 flex flex-col gap-5 items-start justify-start">
                <h2 class="font-semibold text-3xl text-primary font-arial">ჩვენ გვიყვარს, რასაც ვაკეთებთ</h2>
                <p class="w-full text-primary ">
                ჩვენი განვითარების, წარმოების და გაყიდვების საქმიანობა მუდმივად ფართოვდება 1946 წლიდან. ზრდა მოითხოვს მოდერნიზაციას და სტრუქტურას ყველა დონეზე. ჩვენ ამას წარმატებით ვახორციელებთ და ყოველდღიურად ვსწავლობთ ახალ ნივთებს ჩვენი მაღალი ხარისხის სტანდარტების უზრუნველსაყოფად.
                </p>
            </div>
        </div>`,
	},
}

func Populate() {
	Language()

	for _, row := range Interfaces {
		storage.DB.Create(&row)

		storage.DB.
			Model(&row).
			Association("Products").
			Append(
			// repository.FindByID[models.Products](1),
			)

		storage.DB.
			Model(&row).
			Association("News").
			Append(
			// repository.FindByID[models.Posts](4),
			// repository.FindByID[models.Posts](3),
			// repository.FindByID[models.Posts](2),
			// repository.FindByID[models.Posts](1),
			)

		storage.DB.
			Model(&row).
			Association("Events").
			Append(
			// repository.FindByID[models.Posts](12),
			// repository.FindByID[models.Posts](13),
			// repository.FindByID[models.Posts](14),
			// repository.FindByID[models.Posts](15),
			)

		storage.DB.
			Model(&row).
			Association("NewVideos").
			Append(
			// repository.FindByID[models.Files](14),
			// repository.FindByID[models.Files](14),
			)

	}
}
