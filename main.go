package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		panic(err)
	}
	updateConfigs := tgbotapi.NewUpdate(0)
	updateConfigs.Timeout = 0
	updates := bot.GetUpdatesChan(updateConfigs)

	/**************************************************************/
	AddCommand("help", help)
	AddCommand("cat", cat)
	/**************************************************************/

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if string(update.Message.Text[0]) != PREFIX {
			continue
		}
		cmdAction, _ := ParseCommand(update.Message.Text)
		action, ok := actions[cmdAction].(func(*tgbotapi.BotAPI, tgbotapi.Update))
		if !ok {
			continue
		}
		action(bot, update)
	}
}

