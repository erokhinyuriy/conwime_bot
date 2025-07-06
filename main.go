package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"

	cmd "conwime/bot/commands"

	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {

	fmt.Println("Telegram Conwime bot")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TG_TOKEN")

	bot, err := telego.NewBot(token, telego.WithDefaultDebugLogger())

	if err != nil {
		log.Fatal("Cannot create a bot")
		os.Exit(1)
	}

	ctx := context.Background()

	updates, _ := bot.UpdatesViaLongPolling(ctx, nil)

	bh, _ := th.NewBotHandler(bot, updates)

	defer func() { _ = bh.Stop() }()

	initCommands(bh)

	bh.Start()

	checkUpdates(updates, bot, ctx)
}

func initCommands(bh *th.BotHandler) {
	bh.Handle(cmd.ProcessStart, th.CommandEqual("start"))
	bh.Handle(cmd.ProcessHelp, th.CommandEqual("help"))
	bh.Handle(cmd.ProcessAnyMessages, th.AnyMessage())
}

func checkUpdates(updates <-chan telego.Update, bot *telego.Bot, ctx context.Context) {
	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID

			sentMessage, _ := bot.SendMessage(ctx,
				tu.Message(
					tu.ID(chatID),
					update.Message.Text,
				),
			)

			fmt.Printf("Sent Message: %v\n", sentMessage)
		}
	}
}
