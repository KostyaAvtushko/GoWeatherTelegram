package RunBot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"goWeatherTG/pkg/initBot"
	"goWeatherTG/pkg/weatherRequest"
)

func Run(botToken string) {
	bot := initBot.InitBot(botToken)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			messageRecievedText := update.Message.Text
			weather := weatherRequest.GetWeatherCity(messageRecievedText)
			if weather.Status == "ok" {
				messageText := "Temperature: " + fmt.Sprint(weather.Temp) + "℃" + "\n" +
					"Feels like: " + fmt.Sprint(weather.RealTemp) + "℃" + "\n" +
					"Weather Status: " + weather.WeatherStatus + "(" + weather.DescWeatherStatus + ")" + "\n" +
					"Wind: " + fmt.Sprint(weather.WindSpeed) + "mps"

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "No city found")
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
		}
	}
}
