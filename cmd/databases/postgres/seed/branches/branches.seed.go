package branches

import (
	"main/internal/models"
	"main/internal/storage"
)

var Seed = []models.Branches{
	{
		Name:        "ავტოშოპი",
		PhoneNumber: "+995 598 731 127",
		DistrictID:  1,
		Map:         "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d5956.903432826521!2d44.767317978984444!3d41.71077051496979!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x4044733282ac5f3d%3A0x143c0fade1026bbf!2sUGT!5e0!3m2!1sen!2sge!4v1714676710692!5m2!1sen!2sge",
	},
	{
		Name:        "ზეთების მაღაზია",
		PhoneNumber: "+995 570 456 221",
		DistrictID:  2,
		Map:         "https://www.google.com/maps/embed?pb=!1m14!1m12!1m3!1d1252.5022918705854!2d44.866541917288544!3d41.69959899895436!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!5e0!3m2!1sen!2sge!4v1714676162741!5m2!1sen!2sge",
	},
	{
		Name:        "ავტოშოპი",
		PhoneNumber: "+995 594 832 126",
		DistrictID:  1,
		Map:         "https://www.google.com/maps/embed?pb=!1m14!1m12!1m3!1d1252.5022918705854!2d44.866541917288544!3d41.69959899895436!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!5e0!3m2!1sen!2sge!4v1714676162741!5m2!1sen!2sge",
	},
	{
		Name:        "ავტოშოპი",
		PhoneNumber: "+995 592 224 552",
		DistrictID:  1,
		Map:         "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d5956.903432826521!2d44.767317978984444!3d41.71077051496979!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x4044733282ac5f3d%3A0x143c0fade1026bbf!2sUGT!5e0!3m2!1sen!2sge!4v1714676710692!5m2!1sen!2sge",
	},
	{
		Name:        "ნაწილები და ზეთები",
		PhoneNumber: "+995 598 111 422",
		DistrictID:  3,
		Map:         "https://www.google.com/maps/embed?pb=!1m14!1m12!1m3!1d1252.5022918705854!2d44.866541917288544!3d41.69959899895436!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!5e0!3m2!1sen!2sge!4v1714676162741!5m2!1sen!2sge",
	},
	{
		Name:        "ავტომანია",
		PhoneNumber: "+995 558 112 332",
		DistrictID:  1,
		Map:         "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d5956.903432826521!2d44.767317978984444!3d41.71077051496979!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x4044733282ac5f3d%3A0x143c0fade1026bbf!2sUGT!5e0!3m2!1sen!2sge!4v1714676710692!5m2!1sen!2sge",
	},
}

func Populate() {
	for _, row := range Seed {
		storage.DB.Create(&row)
	}
}
