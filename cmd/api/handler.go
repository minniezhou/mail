package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"mail-service/cmd/mailhandle"
)

type Handler struct {
	router *chi.Mux
}

func (c *Config) NewHandler() *Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders:   []string{"Content-Type", "Ahuthorization", "Accept", "X-CSRF-Token"},
		MaxAge:           300,
		AllowCredentials: false,
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("ping"))

	r.Get("/", c.hello)
	r.Post("/send", c.sendMail)

	return &Handler{
		router: r,
	}
}

func (c *Config) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ping mail Service")
	payload := jsonRequest{
		Error:   false,
		Message: "Ping Mail Service",
	}
	writeJson(w, http.StatusAccepted, payload)
}

func (c *Config) sendMail(w http.ResponseWriter, r *http.Request) {
	var mailInfo mailhandle.MailInfo
	err := readJson(w, r, &mailInfo)
	if err != nil {
		fmt.Println(err.Error())
		errorJson(w, "Read email request error")
		return
	}
	err = c.mail.SendMail(mailInfo)
	if err != nil {
		fmt.Println(err.Error())
		errorJson(w, "Send Email Error")
		return
	}

	payload := jsonRequest{
		Error:   false,
		Message: fmt.Sprintf("Send Email Succesfully to %s", mailInfo.To),
	}
	writeJson(w, http.StatusAccepted, payload)
}
