package services

import (
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

type MailService interface {
	SendMail(message string, subject string, email string) interface{}
	GenerateIntruderNotificationMessage(rfid string) string
	// GenerateVerifyAccountOTPMessage(otp string) string
}

type mailService struct {
}

func MailServiceImp() MailService {
	return &mailService{}
}

func (service *mailService) SendMail(message string, subject string, email string) interface{} {

	// err := godotenv.Load()
	// if err != nil {
	// 	panic("Failed to load env file")

	// }

	Username := "qtsolutionservice@gmail.com" //"noreplybleetlimited@gmail.com" //"qtsolutionservice@gmail.com" //"noreplybleetlimited@gmail.com" // os.Getenv("MAIL_SENDER")
	Password := "snhjqvvgqzabtcyq"            //"jadxwdvvgrafavwb"              //"snhjqvvgqzabtcyq"            //"jadxwdvvgrafavwb"              //os.Getenv("MAIL_PASSWORD") //"jadxwdvvgrafavwb"
	Host := "smtp.gmail.com"
	Port := 465 // os.Getenv("MAIL_PORT")) / 587 /

	log.Println("mail env loaded")

	m := gomail.NewMessage()
	m.SetHeader("From", Username)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", message)

	d := gomail.NewDialer(Host, Port, Username, Password)

	// Send the mail
	if err := d.DialAndSend(m); err != nil {
		log.Printf("%v", err)
		return nil
	}

	return email
}

func (service *mailService) GenerateIntruderNotificationMessage(rfid string) string {
	msg := fmt.Sprintf("Dear Admin, \n An intruder with the card id  %s has attempted to gain access", rfid)
	return msg
}

// func (service *mailService) GenerateVerifyAccountOTPMessage(otp string) string {
// 	msg := fmt.Sprintf("kindly use the OTP below to verify your Account, the OTP expires in 2 minutes \n OTP : %s ", otp)
// 	return msg
// }

// func (service *mailService) GenerateWelcomeMessage(otp string) string {
// 	msg := fmt.SprintF(
// 	msg := fmt.Sprintf("Thank you for signing up with us, we assure you a quality and reliable service, have a nice day ", otp)
// 	return msg
// }
