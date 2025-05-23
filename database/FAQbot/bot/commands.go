package bot

type CreateBotCommand struct {
	Name     string `json:"name"`
	Telegram string `json:"telegram"`
	Facebook string `json:"facebook"`
	WhatsApp string `json:"whats_app"`
}

type GetBotByIdCommand struct {
	Id string `json:"id"`
}

type ListBotCommand struct {
}
