package main

import (
	"bufio"
	"net/smtp"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	f, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0666)
	cmd := exec.Command("top", "-b", "-n", "1")
	cmd.Stdin = os.Stdin
	cmd.Stdout = f

	if err != nil {
		panic(err)
	}

	cmd.Run()
	f.Close()

	file, err := os.OpenFile("output.txt", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	line_count := 0
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
		line_count++
	}
	for i := 7; i < line_count; i++ {

		line := lines[i]
		words := strings.Fields(line)

		p_id, err := strconv.Atoi(words[0])
		cpu, err2 := strconv.ParseFloat(words[8], 64)

		if err != nil {
			panic(err)
		}
		if err2 != nil {
			panic(err2)
		}

		data := Data{p_id, cpu, words[11]}
		result := data.cpu_check(cpu)
		if result == false {
			data.sendMail()
		}
	}

}

type Data struct {
	p_id    int
	cpu     float64
	command string
}

var threshold float64 = 10.0

func (d Data) cpu_check(cpu float64) bool {
	if d.cpu > threshold {
		return false
	}
	return true
}

func (d Data) sendMail() string {
	from := "email"
	password := "password"

	toEmailAddress := "email"
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Subject: Warning! By processor\r\n"
	body := "please check... your process id "
	/*
		body2 := " command "
		body3 := " take  "
		body4 := " percent of cpu. hurry up! and resolve it and reply to this message"

	*/
	message := []byte(subject + "\r\n" + body)
	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}
	return "check your mail"
}
