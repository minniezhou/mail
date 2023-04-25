package main

import (
	"fmt"
	"log"
	"net/http"

	"mail-service/cmd/mailhandle"
)

type Config struct {
	mail *mailhandle.MailHandler
}

const (
	webPort = 54321
)

func main() {
	fmt.Println("This is mailing service")
	m := mailhandle.NewMailHandler()
	c := Config{mail: m}
	h := c.NewHandler()
	err := http.ListenAndServe(fmt.Sprintf(":%d", webPort), h.router)

	if err != nil {
		log.Panic(err)
	}
}
