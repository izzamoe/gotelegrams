package query

import (
	"hehe/db"
	"hehe/entity"
)

// RegisterMember is a query that registers a member
func RegisterMember(userData *entity.User) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}

	db.Db.Create(userData)
	return nil

}

// get user by chatid
func GetUserByChatID(chatid string) (*entity.User, error) {
	db, err := db.GetDB()
	if err != nil {
		return nil, err
	}

	var user entity.User
	db.Db.Where("chat_id = ?", chatid).First(&user)
	// check if user is empty
	if user.ID == 0 {
		return nil, nil
	}

	return &user, nil
}
