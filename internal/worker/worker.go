package worker

import "github.com/alhaos/telegramBot/internal/telegram"

type Worker struct {
	telegram *telegram.Telegram
}

func NewWorker(telegram *telegram.Telegram) *Worker {
	return &Worker{telegram: telegram}
}

func (w *Worker) SendMessage(message string) error {
	return w.telegram.Send(message, "@alhaosChan")
}
