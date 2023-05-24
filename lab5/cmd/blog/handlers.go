package main

import (
	"html/template"
	"log"
	"net/http"
)

type featuredPostData struct {
	Title       string
	Subtitle    string
	ImgModifier string
	Author      string
	AuthorImg   string
	PublishDate string
	Category    string
}

type mostRecentPostData struct {
	Title       string
	Subtitle    string
	Image       string
	Author      string
	AuthorImg   string
	PublishDate string
}

func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
	if err != nil {
		// чем отличаются эти ошибки ?
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := struct {
		FeaturedPosts   []featuredPostData
		MostRecentPosts []mostRecentPostData
	}{
		FeaturedPosts:   featuredPosts(),
		MostRecentPosts: mostRecentPosts(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
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

func featuredPosts() []featuredPostData {
	return []featuredPostData{
		{
			Title:       "The Road Ahead",
			Subtitle:    "The road ahead might be paved - it might not be.",
			ImgModifier: "featured-post_first",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/Mat_Vogels.png",
			PublishDate: "September 25, 2015",
		},
		{
			Title:       "From Top Down",
			Subtitle:    "Once a year, go someplace you’ve never been before.",
			ImgModifier: "featured-post_second",
			Author:      "William Wong",
			AuthorImg:   "static/img/William_Wong.png",
			PublishDate: "September 25, 2015",
			Category:    "Adventure",
		},
	}
}

func mostRecentPosts() []mostRecentPostData {
	return []mostRecentPostData{
		{
			Title:       "Still Standing Tall",
			Subtitle:    "Life begins at the end of your comfort zone.",
			Image:       "static/img/Still_Standing_Tall.jpg",
			Author:      "William Wong",
			AuthorImg:   "static/img/William_Wong.png",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Sunny Side Up",
			Subtitle:    "No place is ever as bad as they tell you it’s going to be.",
			Image:       "static/img/Sunny_Side_Up.jpg",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/Mat_Vogels.png",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Water Falls",
			Subtitle:    "We travel not to escape life, but for life not to escape us.",
			Image:       "static/img/Water_Falls.jpg",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/Mat_Vogels.png",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Through the Mist",
			Subtitle:    "Travel makes you see what a tiny place you occupy in the world.",
			Image:       "static/img/Through_the_Mist.jpg",
			Author:      "William Wong",
			AuthorImg:   "static/img/William_Wong.png",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Awaken Early",
			Subtitle:    "Not all those who wander are lost.",
			Image:       "static/img/Awaken_Early.jpg",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/Mat_Vogels.png",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Try it Always",
			Subtitle:    "The world is a book, and those who do not travel read only one page.",
			Image:       "static/img/Try_it_Always.jpg",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/Mat_Vogels.png",
			PublishDate: "9/25/2015",
		},
	}
}
