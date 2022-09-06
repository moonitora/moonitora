package email

import (
	"fmt"
	"github.com/victorbetoni/moonitora/model"
	"gopkg.in/gomail.v2"
	"os"
	"strings"
)

func NotifyMonitor(monitor string, monitoria model.Monitoria) error {

	email := strings.TrimSpace(os.Getenv("APP_EMAIL"))
	user := strings.TrimSpace(os.Getenv("APP_EMAIL_USER"))
	pw := strings.TrimSpace(os.Getenv("APP_EMAIL_PASSWORD"))
	mn := strings.TrimSpace(monitor)

	fmt.Println("From:", email)
	fmt.Println("User:", user)
	fmt.Println("Pw:", pw)
	fmt.Println("Monitor:", mn)

	msg := gomail.NewMessage()
	msg.SetHeader("From", email)
	msg.SetHeader("To", mn)
	msg.SetHeader("Subject", "Monitoria marcada para você")
	msg.SetBody("text/html", "<b>Uma nova monitoria foi marcada para você</b>")

	n := gomail.NewDialer("smtp.gmail.com", 587, user, pw)

	if err := n.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}
