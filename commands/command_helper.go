package commands

import (
	"fmt"
	"log"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ProcessStart(ctx *th.Context, updates telego.Update) error {
	chatID := updates.Message.Chat.ID

	keyboard := tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton("Помощь"), // нужны инлайн кнопки с колбэком
		),
		tu.KeyboardRow(
			tu.KeyboardButton("Отправить локацию").WithRequestLocation(),
			tu.KeyboardButton("Отправить контакт").WithRequestContact(),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите чего-нибудь..")

	message := tu.Message(
		tu.ID(chatID),
		"Привет! Чем могу помочь?",
	).WithReplyMarkup(keyboard).WithProtectContent()

	_, _ = ctx.Bot().SendMessage(ctx, message)

	fmt.Printf("Отправленное сообщение: %v\n", message)
	log.Printf("Отправленное сообщение: %v\n", message)

	return nil
}

func ProcessHelp(ctx *th.Context, updates telego.Update) error {
	chatID := updates.Message.Chat.ID

	_, _ = ctx.Bot().SendMessage(ctx,
		tu.Message(
			tu.ID(chatID),
			"Помощь уже близка!",
		),
	)

	return nil
}

func ProcessAnyMessages(ctx *th.Context, updates telego.Update) error {
	chatID := updates.Message.Chat.ID

	_, _ = ctx.Bot().SendMessage(ctx,
		tu.Message(
			tu.ID(chatID),
			"Упс.. давайте по-конкретнее",
		),
	)

	return nil
}
