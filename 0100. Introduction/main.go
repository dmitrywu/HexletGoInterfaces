package main

import "fmt"

type Notifier interface {
	Notify (message string)	
}

type EmailSender struct {}

func (e EmailSender) Notify(message string){
	fmt.Println("Отправка email:", message)
}

func Send(n Notifier) {
n.Notify("Добро пожаловать")
}

func main() {
	fmt.Print("\033[H\033[2J")

	Send(EmailSender{})


}