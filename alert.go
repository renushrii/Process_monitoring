package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

type Data struct {
	pId     int
	cpu     float64
	command string
}

func NewData(pID int, cpu float64, command string) Data {
	return Data{
		pId:     pID,
		cpu:     cpu,
		command: command,
	}
}

type EmailSender interface {
	sendMails(dataList []Data) error
}

type SmtpEmailSender struct{}

func (d Data) isCPUUsageOverThreshold(threshold float64) bool {
	if d.cpu > threshold {
		return true
	}
	return false
}

func (sms SmtpEmailSender) sendMails(dataList []Data) error {
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")

	toEmailAddress := os.Getenv("TO_EMAIL")
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	auth := smtp.PlainAuth("", from, password, host)

	for _, data := range dataList {
		subject := "Subject: Warning! CPU Usage Over Threshold\r\n"
		body := fmt.Sprintf("Please check process with ID %d and command '%s'. It's using %.2f%% of the CPU.", data.pId, data.command, data.cpu)

		message := []byte(subject + "\r\n" + body)

		err := smtp.SendMail(address, auth, from, to, message)
		if err != nil {
			log.Printf("Failed to send email: %v", err)
			continue
		}
	}

	return nil
}
