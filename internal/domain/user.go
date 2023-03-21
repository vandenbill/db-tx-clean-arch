package domain

type User struct {
	ID       int    `json:"id" gorm:"primarykey"`
	Username string `json:"username" gorm:"unique"`
}
