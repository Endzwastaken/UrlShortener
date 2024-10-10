package memstorage

import (
	"errors"
)

// структура типа хранилища в памяти
type urlsMap struct {
	urls map[string]string
}

// конструктор
func New() *urlsMap {
	return &urlsMap{
		urls: make(map[string]string),
	}
}

// функция получения из мапы
func (u *urlsMap) Get(shortLink string) (string, error) {
	link, exists := u.urls[shortLink]
	if !exists {
		return link, errors.New("there is no link for that short link")
	}
	return link, nil
}

// функция вставки в мапу
func (u *urlsMap) Insert(shortLink string, link string) error {
	u.urls[shortLink] = link
	return nil
}
