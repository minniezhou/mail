package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"mail-service/cmd/mailhandle"
)

type Config struct {
	mail *mailhandle.MailHandler
}

func main() {
	fmt.Println("This is mailing service")
	m := mailhandle.NewMailHandler()
	c := Config{mail: m}
	h := c.NewHandler()
	webPort, _ := strconv.ParseInt(os.Getenv("WEB_PORT_DEFAULT"), 10, 64)
	value, exists := os.LookupEnv("WEB_PORT")
	if exists {
		webPort, _ = strconv.ParseInt(value, 10, 64)
	}
	err := http.ListenAndServe(fmt.Sprintf(":%d", webPort), h.router)

	if err != nil {
		log.Panic(err)
	}
}
