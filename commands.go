package main

import (
	"encoding/json"
	"io/ioutil"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func help(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	message := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.From.FirstName)
	bot.Send(message)
}

func cat(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	httpResp, err := HttpGet("https://api.thecatapi.com/v1/images/search")
	if err != nil {
		return
	}
	response, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}
	var data interface{}
	if err := json.Unmarshal(response, &data); err != nil {
		return
	}
	// You don't need to understand this shit, baka!
	catUrl := data.([]interface{})[0].(map[string]interface{})["url"].(string)
	mediaGroup := []interface{}{
		tgbotapi.NewInputMediaPhoto(tgbotapi.FileURL(catUrl)),
	}
	message := tgbotapi.NewMediaGroup(update.Message.Chat.ID, mediaGroup)
	bot.Send(message)
}
