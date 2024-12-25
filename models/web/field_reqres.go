package web

type FieldReqRes struct {
	Name  string `json:"Name" form:"Name" validate:"required"`
	Type  string `json:"Type" form:"Type" validate:"required"`
	Price uint   `json:"Price" form:"Price" validate:"required"`
}
