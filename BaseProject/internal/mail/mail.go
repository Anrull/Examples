package mail

import (
	"Examples/BaseProject/internal/config"
	"crypto/tls"
	"errors"
	"fmt"
	"net/smtp"
	"net/textproto"
	"strconv"
)

var ErrNoHost = errors.New("no host specified")
var cfg *SMTP

func New(cfg_ *config.Config) {
	port, err := strconv.Atoi(cfg_.SMTP.Port)
	if err != nil {
		panic(err)
	}
	cfg = &SMTP{
		Host:     cfg_.SMTP.Host,
		Port:     port,
		Username: cfg_.SMTP.Username,
		Password: cfg_.SMTP.Password,
	}
}

type SMTP struct {
	Host     string
	Port     int
	Username string
	Password string
}

func Send(to, subject, body string) error {
	if cfg.Host == "" {
		return ErrNoHost
	}

	from := "no-reply@" + cfg.Host
	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
	toList := []string{to}

	header := make(textproto.MIMEHeader)
	header.Set("Content-Type", "text/plain; charset=utf-8")
	header.Set("Subject", subject)
	header.Set("To", to)
	header.Set("From", from)

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	conn, err := tls.Dial("tcp", addr, &tls.Config{ServerName: cfg.Host})
	if err != nil {
		return err
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, cfg.Host)
	if err != nil {
		return err
	}
	defer client.Quit()

	if err = client.Auth(auth); err != nil {
		return err
	}

	if err = client.Mail(from); err != nil {
		return err
	}

	for _, addr := range toList {
		if err = client.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	if err = w.Close(); err != nil {
		return err
	}

	return client.Quit()
}

