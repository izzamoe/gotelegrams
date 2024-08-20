package middlewares

import (
	"context"
	"hehe/entity"
	"hehe/query"
	"log"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// RegisterMemberMiddleware is a middleware that registers a member
func RegisterMemberMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		result, err := query.GetUserByChatID(strconv.FormatInt(update.Message.Chat.ID, 10))
		if err != nil {
			log.Printf("Error getting user by chat id %d: %s", update.Message.Chat.ID, err)
		}

		if result != nil {
			log.Printf("User already registered: %d", update.Message.Chat.ID)
		} else {
			if update.Message != nil {
				log.Printf("Registering member: %d", update.Message.From.ID)
				err := query.RegisterMember(&entity.User{
					ChatID:   strconv.FormatInt(update.Message.Chat.ID, 10),
					Name:     update.Message.From.FirstName,
					Username: update.Message.From.Username,
					Berapakali: entity.Berapakali{
						BerapaX: 0,
					},
				})
				if err != nil {
					log.Printf("Error registering member %d: %s", update.Message.From.ID, err)
				}

			}
		}
		next(ctx, b, update)
	}
}
