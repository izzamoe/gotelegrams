package main

import (
	"context"
	"fmt"
	"hehe/db"
	"hehe/handlers"
	"hehe/middlewares"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
)

func main() {

	db, err := db.NewDB()
	if err != nil {
		fmt.Println("Failed to connect to database")
		panic(err)
	}
	db.Migrate()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithMiddlewares(middlewares.ShowMessageIDMiddleware, middlewares.RegisterMemberMiddleware, middlewares.CountMessage),
		bot.WithDefaultHandler(handlers.DefaultHandler),
	}

	b, err := bot.New("7045295114:AAH2dZrp5_MRzxkbjfXh1L6BePVOvrgZK_M", opts...)
	if err != nil {
		// TODO: HANDLE ERRORs
		panic(err)
	}

	// b.DeleteWebhook(ctx, nil)
	// hasil, err := b.SetWebhook(ctx, &bot.SetWebhookParams{
	// 	URL: "https://x1.izzamoe.my.id",
	// 	// SecretToken: "secret",
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(hasil)

	// go func() {
	// 	http.ListenAndServe(":2000", b.WebhookHandler())
	// }()

	b.RegisterHandler(bot.HandlerTypeMessageText, "/hello", bot.MatchTypeExact, handlers.HelloHandler)

	b.Start(ctx)
	// b.StartWebhook(ctx)

}
