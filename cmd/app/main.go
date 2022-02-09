package main

import "goWeatherTG/pkg/onUpdate"

const (
	BOT_TOKEN = "<my_token>"
)

func main() {
	onUpdate.OnUpdate(BOT_TOKEN)
}
