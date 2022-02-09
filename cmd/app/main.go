package main

import "goWeatherTG/pkg/RunBot"

const (
	BOT_TOKEN = "<my_token>"
)

func main() {
	RunBot.Run(BOT_TOKEN)
}
