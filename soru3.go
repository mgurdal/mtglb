// subs içerisinde bulunan herkese, notify fonksiyonunu kullanarak,
// dead-lock oluimadan, goroutineler ile mesaj gönderebilen kodu yazın.
// (sync.WaitGroup)
package main

import (
	"fmt"
)

func Notify(to, msg string) {
	fmt.Println("sending message:", msg, "to:", to)
}

func main() {
	subs := []string{
		"user1",
		"user2",
		"user3",
	}
	msg := "hello"
	for _, sub := range subs {
		Notify(sub, msg)
	}
}
