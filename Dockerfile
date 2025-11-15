# -----------------------------
# Этап 1 — сборка бинарника Go
# -----------------------------
FROM golang:1.25-alpine AS builder

# Рабочая директория
WORKDIR /app

# Загружаем зависимости отдельно (кеш)
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем бинарник из папки cmd/qa_service
RUN go build -o server ./cmd/qa_service


# ------------------------------------
# Этап 2 — минимальный образ для запуска
# ------------------------------------
FROM alpine:latest

WORKDIR /root/

# Копируем бинарник
COPY --from=builder /app/server .

# Порт API
EXPOSE 8080

# Запуск приложения
CMD ["./server"]
