package dto

//BookUpdateDTO is used for modeling data when user updating a book's data
type BookUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserId      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

//BookCreateDTO is used for modeling data when user creating new book's data
type BookCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserId      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
