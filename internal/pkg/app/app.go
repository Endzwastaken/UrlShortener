package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Endzwastaken/test-task/internal/app/endpoint"
	"github.com/Endzwastaken/test-task/internal/app/service"
	"github.com/gorilla/mux"
)

type App struct {
	e *endpoint.Endpoint
	s *service.Service
}

func New() (*App, error) {
	// создаём флаг "d". Если введён - true - работа с бд, нет - работа с памятью.
	dbflag := flag.Bool("d", false, "used to on/off database")
	// получаем значение из командной строки
	flag.Parse()

	a := &App{}

	a.s = service.New(*dbflag)

	a.e = endpoint.New(a.s)

	// создаём роуты для того, чтобы обработчики знали на что реагировать
	r := mux.NewRouter()
	r.HandleFunc("/", a.e.Form)
	r.HandleFunc("/s", a.e.Shorting)
	r.HandleFunc("/{shortKey}", a.e.Redirect)
	http.Handle("/", r)

	return a, nil

}

func (a *App) Run() error {
	// сообщение о запуске сервера
	fmt.Println("URL Shortener is running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
