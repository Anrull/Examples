package bot

import (
	"Examples/BaseProject/pkg/env"
	"os"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

var Bot *tgbotapi.BotAPI
var Logger *logrus.Logger

func New(log *logrus.Logger) {
	var err error
	Bot, err = tgbotapi.NewBotAPI(env.GetValue("TOKEN_BOT"))
	Logger = log

	if err != nil {
		Logger.Info("error creating new bot")
	}
}

func SendText(txt string) (tgbotapi.Message, error) {
	chatId, err := strconv.Atoi(env.GetValue("TG_CHAT_ID"))
	if err != nil {
		Logger.Info("uncorrect telegram chat id")
		return tgbotapi.Message{}, err
	}
	msg := tgbotapi.NewMessage(int64(chatId), txt)
	return Bot.Send(msg)
}

func Send(message tgbotapi.Chattable) {
	_, err := Bot.Send(message)

	if err != nil {
		Logger.Info("error sending message: ", err)
	}
}

func SendFile(filename, title, Caption string) error {
	fileReader, _ := os.Open(filename)
	defer fileReader.Close()

	inputFile := tgbotapi.FileReader{
		Name:   title,
		Reader: fileReader,
	}

	chatId, err := strconv.Atoi(env.GetValue("TG_CHAT_ID"))
	if err != nil {
		Logger.Info("uncorrect telegram chat id")
		return err
	}

	msg := tgbotapi.NewDocument(int64(chatId), inputFile)
	if Caption != "time" {
		msg.Caption = Caption
	} else {
		msg.Caption = time.Now().Format("2006-01-02 15:04:05")
	}

	_, err = Bot.Send(msg)

	if err != nil {
		Logger.Info("error sending file: ", err)
		return err
	}

	return nil
}
