package main

import (
	"fmt"
)

type Sender interface {
	send()
}

type Exporter interface {
	export()
}

type SmsSender struct {
}

type EmailSender struct {
}

func (s *EmailSender) send() {
	fmt.Println("Sending email...")
}

func (s *SmsSender) send() {
	fmt.Println("Sending sms...")
}

func (s *SmsSender) export() {
	fmt.Println("Export...")
}

func sendAndExport(s Sender) {
	s.send()
	if s, ok := s.(Exporter); ok {
		s.export()
	}
}

func main() {
	s := &SmsSender{}
	var b Sender = s
	a := b.(*SmsSender)
	fmt.Printf("%T\n", a)
	sendAndExport(s)
}
