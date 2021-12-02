package dto

//RegisterDTO is used for modeling data when client doing register
type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min:1"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password form:"password" binding:"required" validate:"min:6"`
}
