package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

type Alerter interface {
	Alert(processMetric ProcessMetric)
}

type EmailAlerter struct {
	From     string
	Password string
	To       []string
	Host     string
	Port     string
}

func NewEmailAlerter() *EmailAlerter {
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	toEmailAddress := os.Getenv("TO_EMAIL")
	to := []string{toEmailAddress}
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")

	return &EmailAlerter{
		From:     from,
		Password: password,
		To:       to,
		Host:     host,
		Port:     port,
	}
}

func (ea *EmailAlerter) Alert(processMetric ProcessMetric) error {
	auth := smtp.PlainAuth("", ea.From, ea.Password, ea.Host+":"+ea.Port)

	subject := "Subject: Warning! CPU Usage Over Threshold\r\n"
	body := fmt.Sprintf("Please check process with ID %d and command '%s'. It's using %.2f%% of the CPU.", processMetric.PID, processMetric.Command, processMetric.CPU)

	message := []byte(subject + "\r\n" + body)

	err := smtp.SendMail(ea.Host+":"+ea.Port, auth, ea.From, ea.To, message)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}

	return nil
}
