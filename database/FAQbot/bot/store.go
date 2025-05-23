package bot

type BotStore interface {
	Create(bot *Bot) (*Bot, error)
	GetById(id string) (*Bot, error)
	List() ([]Bot, error)
}
