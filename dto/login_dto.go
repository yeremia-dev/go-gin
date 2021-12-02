package dto

//LoginDTO is used for modeling data when client doing login
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password form:"password" binding:"required" validate:"min:6"`
}
