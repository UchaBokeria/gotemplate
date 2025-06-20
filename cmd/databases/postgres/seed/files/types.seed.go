package files

import (
	"main/internal/models"
	"main/internal/storage"
)

var Type = []models.File_types{
	{
		Name:     "Archives/Rar",
		Ext:      "rar",
		Max_size: 10 * 1024 * 1024, // 10MB
	},
	{
		Name:     "Image/Jpeg",
		Ext:      "jpeg",
		Max_size: 10 * 1024 * 1024, // 10MB
	},
	{
		Name:     "Image/Jpg",
		Ext:      "jpg",
		Max_size: 10 * 1024 * 1024, // 10MB
	},
	{
		Name:     "Image/Png",
		Ext:      "png",
		Max_size: 10 * 1024 * 1024, // 10MB
	},
	{
		Name:     "Image/Gif",
		Ext:      "gif",
		Max_size: 10 * 1024 * 1024, // 10MB
	},
	{
		Name:     "Video/Mov",
		Ext:      "mov",
		Max_size: 200 * 1024 * 1024, // 200MB
	},
	{
		Name:     "Video",
		Ext:      "mp4",
		Max_size: 100 * 1024 * 1024, // 100MB
	},
	{
		Name:     "PDF",
		Ext:      "pdf",
		Max_size: 20 * 1024 * 1024, // 20MB
	},
	{
		Name:     "Text Document",
		Ext:      "txt",
		Max_size: 10 * 1024 * 1024, // 10MB
	},
	{
		Name:     "Document",
		Ext:      "docx",
		Max_size: 10 * 1024 * 1024, // 10MB
	},
	{
		Name:     "Spreadsheet",
		Ext:      "xlsx",
		Max_size: 10 * 1024 * 1024, // 10MB
	},
	{
		Name:     "Presentation",
		Ext:      "pptx",
		Max_size: 20 * 1024 * 1024, // 20MB
	},
	{
		Name:     "Audio",
		Ext:      "mp3",
		Max_size: 50 * 1024 * 1024, // 50MB
	},
	{
		Name:     "Executable",
		Ext:      "exe",
		Max_size: 100 * 1024 * 1024, // 100MB
	},
	{
		Name:     "Archive",
		Ext:      "zip",
		Max_size: 100 * 1024 * 1024, // 100MB
	},
	{
		Name:     "Font",
		Ext:      "ttf",
		Max_size: 1 * 1024 * 1024, // 1MB
	},
	{
		Name:     "svg",
		Ext:      "svg",
		Max_size: 20 * 1024 * 1024, // 1MB
	},
}

func Types() {
	for _, row := range Type {
		storage.DB.Create(&row)
	}
}
