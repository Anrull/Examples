package bot

import (
	"Examples/BaseProject/internal/config"
	"Examples/BaseProject/internal/logger"
	"Examples/BaseProject/internal/bot/dispatcher"
	"Examples/BaseProject/pkg/tg"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Bot *tgbotapi.BotAPI


func New(cfg *config.Config, ) {
	tg.New()
	logger.Info("Бот подключен")

	Bot = tg.Bot

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 0
	updates := Bot.GetUpdatesChan(u)

	for update := range updates {
		go dispatcher.Dispatcher(&update)
	}
}
