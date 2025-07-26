package dtos

type CategoryListDto struct {
	Search   string `query:"search"`
	ParentID *uint  `query:"parentId"`
	IsActive *bool  `query:"isActive"`
	Level    *int   `query:"level"`
	Page     int    `query:"page,default:1"`
	Limit    int    `query:"limit,default:100"`
	TreeView bool   `query:"treeView,default:true"`
}
