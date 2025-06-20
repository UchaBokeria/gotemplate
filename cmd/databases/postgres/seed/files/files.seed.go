package files

import (
	"main/internal/models"
	"main/internal/storage"
)

var Seed = []models.Files{
	{
		Name:       "Slide Image",
		Original:   "slide-1.jpg",
		Location:   "local",
		Path:       "/assets/images/slide-1.webp",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     3, // Image type ID
	},
	{
		Name:       "Slide-2 Image",
		Original:   "slide-2.jpg",
		Location:   "local",
		Path:       "/assets/images/slide-2.jpg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     3, // Image type ID
	},
	{
		Name:       "Brothers Image",
		Original:   "brothers.jpg",
		Location:   "local",
		Path:       "/assets/images/brothers.jpg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     3, // Image type ID
	},
	{
		Name:       "Example Image",
		Original:   "example.jpg",
		Location:   "local",
		Path:       "/assets/images/example.jpg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     3, // Image type ID
	},
	{
		Name:       "Example Image",
		Original:   "example.jpg",
		Location:   "local",
		Path:       "/assets/images/example.jpg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     3, // Image type ID
	},
	{
		Name:       "Example Image",
		Original:   "example.jpg",
		Location:   "local",
		Path:       "/assets/images/example.jpg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     3, // Image type ID
	},
	{
		Name:       "garantie Image",
		Original:   "garantie.jpg",
		Location:   "local",
		Path:       "/assets/images/garantie.jpg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     3, // Image type ID
	},
	{
		Name:       "mark Image",
		Original:   "mark.jpg",
		Location:   "local",
		Path:       "/assets/images/mark.jpg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     3, // Image type ID
	},
	{
		Name:       "mark Image",
		Original:   "mark.jpg",
		Location:   "local",
		Path:       "/assets/images/mark.jpg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     3, // Image type ID
	},
	{
		Name:       "mark Image",
		Original:   "mark.jpg",
		Location:   "local",
		Path:       "/assets/images/mark.jpg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     3, // Image type ID
	},
	{
		Name:       "mark Image",
		Original:   "mark.jpg",
		Location:   "local",
		Path:       "/assets/images/mark.jpg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     3, // Image type ID
	},
	{
		Name:       "mark Image",
		Original:   "mark.jpg",
		Location:   "local",
		Path:       "/assets/images/mark.jpg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     3, // Image type ID
	},
	{
		Name:       "Example Image",
		Original:   "example.jpg",
		Location:   "local",
		Path:       "/assets/images/example.jpg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     3, // Image type ID
	},
	{
		Name:       "alfashop",
		Original:   "alfashop",
		Location:   "local",
		Path:       "/assets/videos/alfashop.mp4",
		Size:       1024 * 1024, // 1MB
		Base64:     "",
		Compressed: true,
		TypeID:     7, // Video type ID
	},
	{
		Name:       "Category image",
		Original:   "category-1.png",
		Location:   "local",
		Path:       "/assets/images/categories/category-1.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "Category image",
		Original:   "category-2.png",
		Location:   "local",
		Path:       "/assets/images/categories/category-2.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "Category image",
		Original:   "category-3.png",
		Location:   "local",
		Path:       "/assets/images/categories/category-3.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "Category image",
		Original:   "category-4.png",
		Location:   "local",
		Path:       "/assets/images/categories/category-4.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "Category image",
		Original:   "category-5.png",
		Location:   "local",
		Path:       "/assets/images/categories/category-5.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "Category image",
		Original:   "category-6.png",
		Location:   "local",
		Path:       "/assets/images/categories/category-6.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "Category image",
		Original:   "category-7.png",
		Location:   "local",
		Path:       "/assets/images/categories/category-7.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "Category image",
		Original:   "category-8.png",
		Location:   "local",
		Path:       "/assets/images/categories/category-8.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "Category image",
		Original:   "category-9.png",
		Location:   "local",
		Path:       "/assets/images/categories/category-9.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "News image",
		Original:   "news-1.jpeg",
		Location:   "local",
		Path:       "/assets/images/news/news-1.jpeg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     2, // Image type ID
	},
	{
		Name:       "News image",
		Original:   "news-2.jpeg",
		Location:   "local",
		Path:       "/assets/images/news/news-2.jpeg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     2, // Image type ID
	},
	{
		Name:       "News image",
		Original:   "news-3.jpeg",
		Location:   "local",
		Path:       "/assets/images/news/news-3.jpeg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     2, // Image type ID
	},
	{
		Name:       "News image",
		Original:   "news-4.jpeg",
		Location:   "local",
		Path:       "/assets/images/news/news-4.jpeg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     2, // Image type ID
	},
	{
		Name:       "News image",
		Original:   "news-5.png",
		Location:   "local",
		Path:       "/assets/images/news/news-5.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "News image",
		Original:   "news-6.jpeg",
		Location:   "local",
		Path:       "/assets/images/news/news-6.jpeg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     2, // Image type ID
	},
	{
		Name:       "News image",
		Original:   "news-7.jpeg",
		Location:   "local",
		Path:       "/assets/images/news/news-7.jpeg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     2, // Image type ID
	},
	{
		Name:       "News image",
		Original:   "news-8.jpeg",
		Location:   "local",
		Path:       "/assets/images/news/news-8.jpeg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     2, // Image type ID
	},
	{
		Name:       "News image",
		Original:   "news-9.jpeg",
		Location:   "local",
		Path:       "/assets/images/news/news-9.jpeg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     2, // Image type ID
	},
	{
		Name:       "docalfashop",
		Original:   "docalfashop.pdf",
		Location:   "local",
		Path:       "/uploads/docalfashop.pdf",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     8, // Image type ID
	},
	{
		Name:       "my-project-(1)-1-1",
		Original:   "my-project-(1)-1-1",
		Location:   "local",
		Path:       "/assets/images/products/my-project-(1)-1-1.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "my-project-(1)-1-2",
		Original:   "my-project-(1)-1-2",
		Location:   "local",
		Path:       "/assets/images/products/my-project-(1)-1-2.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "my-project-(1)-1-3",
		Original:   "my-project-(1)-1-3",
		Location:   "local",
		Path:       "/assets/images/products/my-project-(1)-1-3.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "my-project-(1)-1-4",
		Original:   "my-project-(1)-1-4",
		Location:   "local",
		Path:       "/assets/images/products/my-project-(1)-1-4.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "my-project-(1)-1-5",
		Original:   "my-project-(1)-1-5",
		Location:   "local",
		Path:       "/assets/images/products/my-project-(1)-1-5.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "my-project-(1)-1-6",
		Original:   "my-project-(1)-1-6",
		Location:   "local",
		Path:       "/assets/images/products/my-project-(1)-1-6.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "my-project-(1)-1-7",
		Original:   "my-project-(1)-1-7",
		Location:   "local",
		Path:       "/assets/images/products/my-project-(1)-1-7.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "my-project-(1)-1-8",
		Original:   "my-project-(1)-1-8",
		Location:   "local",
		Path:       "/assets/images/products/my-project-(1)-1-8.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "my-project-(1)-1-9",
		Original:   "my-project-(1)-1-9",
		Location:   "local",
		Path:       "/assets/images/products/my-project-(1)-1-9.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "my-project-(1)-1-10",
		Original:   "my-project-(1)-1-10",
		Location:   "local",
		Path:       "/assets/images/products/my-project-(1)-1-10.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "my-project-(1)-1-11",
		Original:   "my-project-(1)-1-11",
		Location:   "local",
		Path:       "/assets/images/products/my-project-(1)-1-11.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "ge.svg",
		Original:   "ge.svg",
		Location:   "local",
		Path:       "/assets/icons/languages/ge.svg",
		Size:       2000, // 2 KB
		Base64:     "",
		Compressed: false,
		TypeID:     17, // SVG type ID
	},
	{
		Name:       "en.svg",
		Original:   "en.svg",
		Location:   "local",
		Path:       "/assets/icons/languages/en.svg",
		Size:       2000, // 2 KB
		Base64:     "",
		Compressed: false,
		TypeID:     17, // SVG type ID
	},
	{
		Name:       "ru.svg",
		Original:   "ru.svg",
		Location:   "local",
		Path:       "/assets/icons/languages/ru.svg",
		Size:       2000, // 2 KB
		Base64:     "",
		Compressed: false,
		TypeID:     17, // SVG type ID
	},
}

func Populate() {
	for _, row := range Seed {
		storage.DB.Create(&row)
	}
}
