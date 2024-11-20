package dispatcher

import (
	"Examples/BaseProject/pkg/tg"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Dispatcher(update *tgbotapi.Update) {
	if update.Message != nil {
		if update.Message.Document != nil {
			// ...
			return
		}
		message := update.Message
		if message.IsCommand() {
			// CommandsHandling(message)
		} else {
			// MessageHandler(message)
		}
		return
	}
	if update.CallbackQuery != nil {
		query := update.CallbackQuery
		lstQ := strings.Split(query.Data, ";")

		fmt.Println(lstQ)
		tg.Request(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
	}
}