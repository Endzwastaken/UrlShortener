package dbstorage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DbStorage struct {
}

// конструктор
func New() *DbStorage {
	return &DbStorage{}
}

// функция добавления url в бд
func (d *DbStorage) Insert(shortLink string, link string) error {
	// подключаемся к бд
	connStr := "user=postgres password=admin dbname=task sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// оставляем разрыв соединения для последнего момента при помощи defer
	defer db.Close()

	// вставляем данные
	_, err = db.Exec("insert into public.\"UrlLinks\" (shortlink, originallink) values ($1, $2)",
		shortLink, link)
	if err != nil {
		panic(err)
	}
	return nil
}

// функция получения значения полной ссылке по короткой
func (d *DbStorage) Get(shortLink string) (string, error) {
	connStr := "user=postgres password=admin dbname=task sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// получаем строки, удовлетворяющие запросу
	rows, err := db.Query("select * from public.\"UrlLinks\" where shortlink = $1", shortLink)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// читаем столбцы и берём второй
	var a, b string
	for rows.Next() {
		err := rows.Scan(&a, &b)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return b, nil
}
