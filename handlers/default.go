package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// DefaultHandler is a default handler for the bot
func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	msg := &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Say /hello",
	}

	b.SendMessage(ctx, msg)

}
