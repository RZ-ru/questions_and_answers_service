package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"qa_service/internal/handlers"
	"qa_service/internal/repository"
)

func main() {
	// Загружаем .env
	_ = godotenv.Load()

	cfg := repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	// Подключение к БД
	database, err := repository.Connect(cfg)
	if err != nil {
		log.Fatal("failed to connect to batabase: ", err)
	}

	log.Println("database connected:", database != nil)

	// Создаём репозиторий
	repo := repository.NewRepository(database)

	// Хендлеры
	questionsHandler := handlers.NewQuestionHandler(repo)
	answersHandler := handlers.NewAnswerHandler(repo)

	// Роутер
	r := chi.NewRouter()

	// --- Questions ---
	r.Post("/questions/", questionsHandler.CreateQuestion)
	r.Get("/questions/", questionsHandler.GetAllQuestions)
	r.Get("/questions/{id}", questionsHandler.GetQuestion)
	r.Delete("/questions/{id}", questionsHandler.DeleteQuestion)

	// --- Answers ---
	r.Post("/questions/{id}/answers/", answersHandler.CreateAnswer)
	r.Get("/answers/{id}", answersHandler.GetAnswer)
	r.Delete("/answers/{id}", answersHandler.DeleteAnswer)

	// Запуск HTTP сервера
	log.Println("server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
