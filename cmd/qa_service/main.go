package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"

	"qa_service/internal/handlers"
	"qa_service/internal/logger"
	"qa_service/internal/repository"
	"qa_service/internal/services"
)

func main() {
	// Инициализируем логгер
	logger.Init()

	// Логируем старт приложения
	log.Info().
		Msg("starting application...")
	// Загружаем переменные окружения из .env файла
	_ = godotenv.Load()
	// Конфигурация подключения к базе данных
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
		log.Fatal().
			Err(err).
			Msg("failed to connect to database")
	}
	// Я думаю понятно
	log.Info().
		Bool("connected", database != nil).
		Msg("database connected")

	// Инициализируем репозитории (доступ к БД)
	repo := repository.NewRepository(database)
	// Инициализируем сервисы (бизнес-логика)
	questionService := services.NewQuestionService(repo.Questions)
	answerService := services.NewAnswerService(repo.Answers, repo.Questions)
	// Инициализируем обработчики HTTP (слой API)
	questionsHandler := handlers.NewQuestionHandler(questionService)
	answersHandler := handlers.NewAnswerHandler(answerService)

	// Роутер
	r := chi.NewRouter()
	// Глобальное логирование всех входящих HTTP-запросов
	r.Use(handlers.LoggingMiddleware)

	// Обработка ендпоинтов вопросов
	r.Post("/questions/", questionsHandler.CreateQuestion)
	r.Get("/questions/", questionsHandler.GetAllQuestions)
	r.Get("/questions/{id}", questionsHandler.GetQuestion)
	r.Delete("/questions/{id}", questionsHandler.DeleteQuestion)

	// Обработка ендпоинтов ответов
	r.Post("/questions/{id}/answers/", answersHandler.CreateAnswer)
	r.Get("/answers/{id}", answersHandler.GetAnswer)
	r.Delete("/answers/{id}", answersHandler.DeleteAnswer)

	// Запуск HTTP сервера
	log.Info().
		Msg("server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to start the server")
	}
}
