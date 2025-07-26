package dtos

type CategoryListDto struct {
	Search      string `query:"search"`
	Path        string `query:"path"`
	Level       *int   `query:"level"`
	ParentID    *uint  `query:"parentId"`
	Type        string `query:"type"`
	VehicleIDs  []uint `query:"vehicleIds"`
	MakerIDs    []uint `query:"makerIds"`
	ModelIDs    []uint `query:"modelIds"`
	TypeIDs     []uint `query:"typeIds"`
	CategoryIDs []uint `query:"categoryIds"`
	IsActive    *bool  `query:"isActive"`
	Page        int    `query:"page,default:1"`
	Limit       int    `query:"limit,default:100"`
	TreeView    bool   `query:"treeView,default:true"`
}
