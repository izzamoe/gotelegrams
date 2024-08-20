package query

import (
	"hehe/db"
	"hehe/entity"
)

func GetBerapakaliByChat(chatID string) (int, error) {
	db, err := db.GetDB()

	if err != nil {
		return 0, err
	}
	var user entity.User
	if err := db.Db.Preload("Berapakali").Where("chat_id = ?", chatID).First(&user).Error; err != nil {
		return 0, err
	}
	return user.Berapakali.BerapaX, nil
}

func CreateBerapakali(berapakali *entity.Berapakali) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}

	db.Db.Create(berapakali)
	return nil
}
func UpdateBerapakali(chatID string) error {
	db, err := db.GetDB()

	if err != nil {
		return err
	}

	var user entity.User
	if err := db.Db.Preload("Berapakali").Where("chat_id = ?", chatID).First(&user).Error; err != nil {
		return err
	}

	user.Berapakali.BerapaX++
	return db.Db.Save(&user.Berapakali).Error
}

// update count berapakali + 1
// func UpdateBerapakali(chatID string) error {
// 	db, err := db.GetDB()
// 	if err != nil {
// 		return err
// 	}

// 	var user entity.User
// 	if err := db.Db.Preload("Berapakali").Where("chat_id = ?", chatID).First(&user).Error; err != nil {
// 		return err
// 	}

// 	user.Berapakali.BerapaX++
// 	db.Db.Save(&user.Berapakali)
// 	return nil
// }

// get Berapakali by chatid
// func GetBerapakaliByChatID(chatid string) (*entity.Berapakali, error) {
// 	db, err := db.GetDB()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var berapakali entity.Berapakali
// 	db.Db.Where("chat_id = ?", chatid).First(&berapakali)
// 	// check if user is empty
// 	if berapakali.ID == 0 {
// 		return nil, nil
// 	}

// 	return &berapakali, nil
// }

// create new berapakali
// func CreateBerapakali(berapakali *entity.User) error {
// 	GetUserByChatID(berapakali.ChatID)

// 	if err != nil {
// 		return err
// 	}

// 	db.Db.Create(berapakali)
// 	return nil
// }

// func CountMessage(chatid *string) error {
// 	db, err := db.GetDB()
// 	if err != nil {
// 		return err
// 	}
// 	db.
// 	// Your code logic here

// 	return nil
// }
