package main

import (
	"errors"
	"fmt"
	"strings"
)

var ErrMessageTooLong = errors.New("message too long")

type Notifier interface {
	Notify(to, msg string) error
}

type EmailNotifier struct {
	From string
}

type SMSNotifier struct {
	SenderID string
}

type FakeNotifier struct {
	Sent []string
}

func (fn *FakeNotifier) Notify(to, msg string) error {
	fn.Sent = append(fn.Sent, fmt.Sprintf("%s : %s", to, msg))
	return nil
}

func (en *EmailNotifier) Notify(to, msg string) error {
	fmt.Printf("email %s -> %s: %s\n", en.From, to, msg)
	return nil
}

func (sm *SMSNotifier) Notify(to, msg string) error {
	msgLen := len(msg)
	const maxSMSBytes = 160
	if msgLen > maxSMSBytes {
		return fmt.Errorf("SMS message to %s -> %d bytes exceeds %d: %w", to, msgLen, maxSMSBytes, ErrMessageTooLong)
	}
	fmt.Printf("SMS %s -> %s: %s\n", sm.SenderID, to, msg)
	return nil
}
func NotifyAll(ns []Notifier, to, msg string) error {
	var errs []error
	for _, n := range ns {
		err := n.Notify(to, msg)

		if err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func main() {
	fake := &FakeNotifier{}
	notifiers := []Notifier{&EmailNotifier{}, &SMSNotifier{}, fake}

	messages := []string{strings.Repeat("a", 200), "hello"}

	for _, message := range messages {
		err := NotifyAll(notifiers, "Jahdiel", message)

		if err != nil {
			fmt.Println(err)
			fmt.Println(errors.Is(err, ErrMessageTooLong))
		}
		fmt.Println("Notify all returned nil")
	}
	fmt.Printf("%q\n", fake.Sent)

}
