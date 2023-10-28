package models

type Room struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"room_id"`
	Name string `json:"room_name"`
	Count int `json:"user_count"`
	Host string `json:"host_user"`
	IsFinished bool `json:"is_finished"`
	CreatedAt string`json:"created_at"`
}