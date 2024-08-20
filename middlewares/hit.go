package middlewares

import (
	"context"
	"hehe/query"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Count berapa kali dia mengirimkan pesan
func CountMessage(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message != nil {
			query.UpdateBerapakali(strconv.FormatInt(update.Message.Chat.ID, 10))
		}
		next(ctx, b, update)
	}
}
