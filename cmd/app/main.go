package main

import "goWeatherTG/pkg/onUpdate"

const (
	BOT_TOKEN = "5166170813:AAEO-TEYhZdUT-PtWx1kIXbu3dpTNGI-VcY"
)

func main() {
	onUpdate.OnUpdate(BOT_TOKEN)
}
