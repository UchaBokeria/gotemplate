package interfaces

import (
	"main/internal/models"
	"main/internal/storage"
)

var Contacts = []models.Interface_contact{
	{
		InterfaceID:    1,
		Ver:            1,
		Name:           "Standard",
		Slug:           "standard",
		Phone:          "( +995 599 ) 156 862",
		Email:          "support@alfashop.ge",
		Location:       "ვარკეთილი, IV მიკრო/რაიონი, 2 რიგი, შენობა №14-ის მოპირდაპირე მხარეს",
		LocationLink:   "https://maps.app.goo.gl/4xdAtSwxpuDroWg7A",
		LocationIframe: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2978.986020935193!2d44.86283857657784!3d41.69923747659316!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x40440d96fddef847%3A0xb243526aef974d8!2sAlfa%20Motors!5e0!3m2!1sen!2sge!4v1714773355601!5m2!1sen!2sge",
		ShortDesc:      "ჩვენ ვავითარებთ, ვაწარმოებთ და ვაწარმოებთ მაღალი ხარისხის საპოხი მასალებს, რომლებიც თანაბრად ცნობილია როგორც გერმანიაში, ასევე საერთაშორისო ბაზარზე. ჩვენს უახლეს, შიდა ლაბორატორიებში, მაღალი ხარისხის პროდუქტები კოორდინირებულია წამყვანი ავტომწარმოებლების მოწონების მისაღებად და მომავალში ახალი სპეციფიკაციების შესასრულებლად. ხარისხის მენეჯმენტი უზრუნველყოფს პროდუქტის მაღალ სტანდარტს, რომელიც გარანტირებულია წარმოებიდან გადაზიდვამდე",
		Copyright:      "AlfaShop ყველა უფლება დაცულია",
	},
}

func Contact() {
	for _, row := range Contacts {
		storage.DB.Create(&row)
	}
}
