package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Импортируем для возможности подключения к MySQL
	"github.com/jmoiron/sqlx"
)

const (
	port         = ":3000"
	dbDriverName = "mysql"
)

func main() {
	db, err := openDB() // Открываем соединение к базе данных в самом начале
	if err != nil {
		log.Fatal(err)
	}

	dbx := sqlx.NewDb(db, dbDriverName) // Расширяем стандартный клиент к базе

	mux := http.NewServeMux() // Сущность Mux, которая позволяет маршрутизировать запросы к определенным обработчикам,
	// зависимости от пути, по которому перешёл пользователь
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("/home", index(dbx)) // Прописываем, что по пути /home выполнится наш index, отдающий нашу страницу
	mux.HandleFunc("/post", post)
	log.Println("Start server at port " + port) // Пишем в консоль о том, что стартуем сервер
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err) // Падаем с логированием ошибки, в случае если не получилось запустить сервер
	}
}

func openDB() (*sql.DB, error) {
	// Здесь прописываем соединение к базе данных
	return sql.Open(dbDriverName, "root:2090870Ma!@tcp(localhost:3000)/blog?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true")
}
