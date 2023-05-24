package main

import (
	"github.com/jmoiron/sqlx"
	"html/template"
	"log"
	"net/http"
)

type indexPageData struct {
	Title           string
	Subtitle        string
	FeaturedPosts   []PostData
	MostRecentPosts []PostData
}

type PostData struct {
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Image       string `db:"image_url"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
}

// IndexView
func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) { // Функция для отдачи страницы

		
			posts, err := featuredPosts(db)
			if err != nil {
				http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
				log.Println(err)
				return // Не забываем завершить выполнение ф-ии
			}
			recentPosts, err := mostRecentPosts(db)
			if err != nil {
				http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
				log.Println(err)
				return // Не забываем завершить выполнение ф-ии
			}

			ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
			if err != nil {
				http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
				log.Println(err)
				return // Не забываем завершить выполнение ф-ии
			}

			data := indexPageData{
				Title:           "Blog for traveling",
				Subtitle:        "My best blog for adventures and burgers",
				FeaturedPosts:   posts,
				MostRecentPosts: recentPosts,
			}

			err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
			if err != nil {
				http.Error(w, "Internal Server Error", 500)
				log.Println(err)
				return
			}
		
		log.Println("Request completed successfully from ", r.Host)
	}
}

func featuredPosts(db *sqlx.DB) ([]PostData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			author,
			image_url,
			author_url,
			publish_date
		FROM
			post
		WHERE featured = 1
	`
	// Составляем SQL-запрос для получения записей для секции featured-posts

	var posts []PostData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return posts, nil
}

func mostRecentPosts(db *sqlx.DB) ([]PostData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			author,
			image_url,
			author_url,
			publish_date
		FROM
			post
		WHERE featured = 0
	` // Составляем SQL-запрос для получения записей для секции featured-posts

	var posts []PostData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return posts, nil
}

func post(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/the-road-ahead.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}
