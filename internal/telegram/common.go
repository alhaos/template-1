package telegram

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

type Telegram struct {
	token string
}

// NewTelegram create new Telegram instance
func NewTelegram(token string) *Telegram {
	return &Telegram{token: token}
}

// Send
func (t *Telegram) Send(text string, chatID string) error {

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.token)

	resp, err := http.PostForm(apiURL, url.Values{
		"chat_id": {chatID},
		"text":    {text},
	})

	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	slog.Info("resp", "body", string(body))

	defer resp.Body.Close()

	return nil
}
