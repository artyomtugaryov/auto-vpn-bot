package tgbotwrap

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func GetUpdatesConfig() tgbotapi.UpdateConfig {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	return updateConfig
}
