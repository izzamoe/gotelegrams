package query

// import (
// 	"hehe/entity"
// 	"testing"

// 	"github.com/glebarez/sqlite"
// 	"gorm.io/gorm"
// )

// // Fungsi untuk mendapatkan BerapaX berdasarkan ChatID
// func UpdateBerapakali(db *gorm.DB, chatID string) error {
// 	var user entity.User
// 	if err := db.Preload("Berapakali").Where("chat_id = ?", chatID).First(&user).Error; err != nil {
// 		return err
// 	}

// 	user.Berapakali.BerapaX++
// 	return db.Save(&user.Berapakali).Error
// }
// func GetBerapakaliByChatID(db *gorm.DB, chatID string) (int, error) {
// 	var user entity.User
// 	if err := db.Preload("Berapakali").Where("chat_id = ?", chatID).First(&user).Error; err != nil {
// 		return 0, err
// 	}
// 	return user.Berapakali.BerapaX, nil
// }

// func TestUpdateBerapakali(t *testing.T) {
// 	// Membuat database SQLite baru dalam memori untuk pengujian
// 	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("failed to connect to database: %v", err)
// 	}

// 	// Migrasi schema database
// 	err = db.AutoMigrate(&entity.User{}, &entity.Berapakali{})
// 	if err != nil {
// 		t.Fatalf("failed to migrate database: %v", err)
// 	}

// 	// Membuat user baru
// 	user := &entity.User{
// 		ChatID:   "123",
// 		Name:     "John Doe",
// 		Username: "johndoe",
// 		Berapakali: entity.Berapakali{
// 			BerapaX: 1,
// 		},
// 	}

// 	// Menyimpan user ke database
// 	db.Create(user)

// 	// Update Berapakali
// 	err = UpdateBerapakali(db, "123")
// 	if err != nil {
// 		t.Fatalf("failed to update Berapakali: %v", err)
// 	}

// 	// Membaca BerapaX setelah update
// 	berapakali, err := GetBerapakaliByChatID(db, "123")
// 	if err != nil {
// 		t.Fatalf("failed to get BerapaX by ChatID after update: %v", err)
// 	}

// 	// Memastikan BerapaX bertambah
// 	if berapakali != 2 {
// 		t.Errorf("expected BerapaX to be 2 after update, got %d", berapakali)
// 	}
// }
