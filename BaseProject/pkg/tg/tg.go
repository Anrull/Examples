package tg

import (
	"Examples/BaseProject/internal/logger"
	"Examples/BaseProject/pkg/env"
	"os"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Bot *tgbotapi.BotAPI

func New() {
	var err error
	Bot, err = tgbotapi.NewBotAPI(env.GetValue("TOKEN_BOT"))

	if err != nil {
		logger.Info("error creating new bot")
	}
}

func SendText(txt string) (tgbotapi.Message, error) {
	chatId, err := strconv.Atoi(env.GetValue("TG_CHAT_ID"))
	if err != nil {
		logger.Info("uncorrect telegram chat id")
		return tgbotapi.Message{}, err
	}
	msg := tgbotapi.NewMessage(int64(chatId), txt)
	return Bot.Send(msg)
}

func Send(message tgbotapi.Chattable) {
	_, err := Bot.Send(message)

	if err != nil {
		logger.Info("error sending message: ", err)
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
		logger.Info("uncorrect telegram chat id")
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
		logger.Info("error sending file: ", err)
		return err
	}

	return nil
}

// Request sends a request to the Bot.
//
// The function takes a parameter of type tgbotapi.Chattable, which represents the request to be sent.
// It returns nothing.
// If there is an error while sending the request, it logs the error message.
func Request(c tgbotapi.Chattable) {
	_, err := Bot.Request(c)
	if err != nil {
		logger.Info("Error sending message:", err)
	}
}