package models

type UserReqModel struct {
	Name string `json:"user_name"`
	Password string `json:"password"`
}