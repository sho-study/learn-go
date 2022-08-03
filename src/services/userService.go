package services

import (
	"my-app/src/database"
	"my-app/src/entities"
	"my-app/src/utils"
)

func CreateUser(name string) entities.User {
	user := entities.User{
		ID:   utils.GetUUID(),
		NAME: name,
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		panic("何かがおかしい")
	}

	return user
}
