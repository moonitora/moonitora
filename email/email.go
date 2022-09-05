package email

import (
	"fmt"
	"github.com/victorbetoni/moonitora/model"
	"gopkg.in/gomail.v2"
	"os"
)

func NotifyMonitor(monitor string, monitoria model.Monitoria) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", os.Getenv("APP_EMAIL"))
	msg.SetHeader("To", monitor)
	msg.SetHeader("Subject", "Monitoria marcada para você")
	msg.SetBody("text/html", "<b>Uma nova monitoria foi marcada para você</b>")

	fmt.Println(monitor)
	fmt.Println(os.Getenv("APP_EMAIL_USER"))
	fmt.Println(os.Getenv("APP_EMAIL_PASSWORD"))

	n := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("APP_EMAIL_USER"), os.Getenv("APP_EMAIL_PASSWORD"))

	if err := n.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}
