package service

import (
	"math/rand"
	"time"

	dbstorage "github.com/Endzwastaken/test-task/pkg/dbStorage"
	memstorage "github.com/Endzwastaken/test-task/pkg/memStorage"
)

type Storage interface {
	Insert(string, string) error
	Get(string) (string, error)
}

type Service struct {
	s Storage
}

func New(dbFlag bool) *Service {
	if dbFlag {
		return &Service{
			s: dbstorage.New(),
		}
	} else {
		return &Service{
			s: memstorage.New(),
		}
	}
}

func (s *Service) Insert(shortLink string, link string) error {
	return s.s.Insert(shortLink, link)
}

func (s *Service) Get(shortLink string) (string, error) {
	return s.s.Get(shortLink)
}

// функция генерации случайного значения
func (s *Service) GenerateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	// делаем генерацию случайной из-за ключа - времени
	rand.NewSource(time.Now().UnixNano())
	// заполняем итоговую коротку ссылку по одному элементу
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}
