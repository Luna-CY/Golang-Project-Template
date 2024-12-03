package request

type BaseListRequest struct {
	Page int `json:"page" validate:"optional" minimum:"1" maximum:"100" binding:"omitempty,min=1,max=100"` // page
	Size int `json:"size" validate:"optional" minimum:"5" maximum:"50" binding:"omitempty,min=5,max=50"`   // number of items per page
}

type BaseBackendListRequest struct {
	Page int `json:"page" validate:"optional" minimum:"1" binding:"omitempty,min=1"`                         // page
	Size int `json:"size" validate:"optional" minimum:"10" maximum:"500" binding:"omitempty,min=10,max=100"` // number of items per page
}
