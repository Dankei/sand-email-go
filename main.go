package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	gomail "gopkg.in/gomail.v2"
)

func main() {

	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
    }
	key := os.Getenv("KEY")
	from := os.Getenv("FROM")
	host := os.Getenv("HOST")
	port,_ := strconv.Atoi(os.Getenv("PORT"))
	to := os.Getenv("FROM")

	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "Hello! Teste de envio de email")
	msg.SetBody("text/html", "<h1>Hello!</h1><br><p>Teste de envio de email</p>")

	n := gomail.NewDialer(host, port, from,key)
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
}