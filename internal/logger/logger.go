package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() {
	// Красивый вывод в консоль
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	// Если хочешь формат JSON — закомментируй строку выше
}
