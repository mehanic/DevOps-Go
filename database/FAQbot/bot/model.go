package bot

type Bot struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Telegram string `json:"telegram"`
	Facebook string `json:"facebook"`
	WhatsApp string `json:"whats_app"`
}
