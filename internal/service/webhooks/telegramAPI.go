package webhooks

import (
	"SignalManager/pkg/logger"
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TelegramApiService wraps Telegram API calls with error logging
// No automatic retries - users can manually retry if needed
type TelegramApiService struct {
	bot *tgbotapi.BotAPI
}

// NewTelegramApiService creates a new Telegram API service instance
func NewTelegramApiService(token string) (*TelegramApiService, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create telegram bot: %w", err)
	}

	return &TelegramApiService{
		bot: bot,
	}, nil
}

// SendMessage sends a message without retry
//
// Returns Telegram API response or error if sending fails
func (s *TelegramApiService) SendMessage(config tgbotapi.MessageConfig) (tgbotapi.Message, error) {
	return s.bot.Send(config)
}

// EditMessageText edits message text without retry
//
// Returns Telegram API response or error if editing fails
func (s *TelegramApiService) EditMessageText(config tgbotapi.EditMessageTextConfig) (tgbotapi.Message, error) {
	return s.bot.Send(config)
}

// EditMessageReplyMarkup edits message reply markup without retry
//
// Returns Telegram API response or error if editing fails
func (s *TelegramApiService) EditMessageReplyMarkup(config tgbotapi.EditMessageReplyMarkupConfig) (tgbotapi.Message, error) {
	return s.bot.Send(config)
}

// DeleteMessage deletes a message without retry
//
// Returns Telegram API response or error if deletion fails
func (s *TelegramApiService) DeleteMessage(config tgbotapi.DeleteMessageConfig) (*tgbotapi.APIResponse, error) {
	return s.bot.Request(config)
}

// DeleteMessagesResult contains statistics about bulk deletion results
type DeleteMessagesResult struct {
	Total   int `json:"total"`
	Deleted int `json:"deleted"`
	Failed  int `json:"failed"`
}

// DeleteMessages deletes multiple messages in bulk
// Silently continues if some messages fail to delete
//
// Returns statistics about deletion results
func (s *TelegramApiService) DeleteMessages(chatID int64, messageIDs []int) DeleteMessagesResult {
	deleted := 0
	failed := 0

	for _, messageID := range messageIDs {
		config := tgbotapi.DeleteMessageConfig{
			ChatID:    chatID,
			MessageID: messageID,
		}

		_, err := s.DeleteMessage(config)
		if err != nil {
			failed++
			logger.Error("[TelegramAPI] Failed to delete message: chat_id=%d, message_id=%d, error=%v",
				chatID, messageID, err)
		} else {
			deleted++
		}

		// Small delay to avoid rate limiting
		time.Sleep(50 * time.Millisecond)
	}

	logger.Info("[TelegramAPI] Bulk message deletion completed: chat_id=%d, total=%d, deleted=%d, failed=%d",
		chatID, len(messageIDs), deleted, failed)

	return DeleteMessagesResult{
		Total:   len(messageIDs),
		Deleted: deleted,
		Failed:  failed,
	}
}

// SendDocument sends a document without retry
//
// Returns Telegram API response or error if sending fails
func (s *TelegramApiService) SendDocument(config tgbotapi.DocumentConfig) (tgbotapi.Message, error) {
	return s.bot.Send(config)
}

// AnswerCallbackQuery answers a callback query without retry
//
// Returns Telegram API response or error if answering fails
func (s *TelegramApiService) AnswerCallbackQuery(config tgbotapi.CallbackConfig) (*tgbotapi.APIResponse, error) {
	return s.bot.Request(config)
}

// SendMessageSafe sends a message with graceful degradation
// Returns nil on failure instead of throwing exception
//
// Returns Telegram API response or nil on failure
func (s *TelegramApiService) SendMessageSafe(config tgbotapi.MessageConfig) *tgbotapi.Message {
	msg, err := s.SendMessage(config)
	if err != nil {
		logger.Error("[TelegramAPI] SendMessageSafe failed: error=%v", err)
		return nil
	}
	return &msg
}

// SetDebug enables or disables debug mode for the bot
func (s *TelegramApiService) SetDebug(debug bool) {
	s.bot.Debug = debug
}

// GetBot returns the underlying bot API instance
func (s *TelegramApiService) GetBot() *tgbotapi.BotAPI {
	return s.bot
}
