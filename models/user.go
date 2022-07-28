package models

type UserModel struct {
	BaseModel
	Name  string `json:"name"`
	Email string `json:"email"`
}
