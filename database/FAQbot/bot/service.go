package bot

type BotService interface {
	CreateBot(cmd *CreateBotCommand) (*Bot, error)
	GetBotById(cmd *GetBotByIdCommand) (*Bot, error)
	ListBot(cmd *ListBotCommand) ([]Bot, error)
}

type botService struct {
	store BotStore
}

func NewBotService(s BotStore) BotService {
	return &botService{store: s}
}

func (b *botService) CreateBot(cmd *CreateBotCommand) (*Bot, error) {
	//check for empty fields
	return b.store.Create(&Bot{
		Name:     cmd.Name,
		Telegram: cmd.Telegram,
		Facebook: cmd.Facebook,
		WhatsApp: cmd.WhatsApp,
	})
}

func (b *botService) GetBotById(cmd *GetBotByIdCommand) (*Bot, error) {
	return b.store.GetById(cmd.Id)
}

func (b *botService) ListBot(cmd *ListBotCommand) ([]Bot, error) {
	return b.store.List()
}
