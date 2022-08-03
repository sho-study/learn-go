package entities

type User struct {
	ID   string `gorm:"primary_key"`
	NAME string
}
