package scenarious


type Scenarios struct {
	Id string
	BotId      string
	Name       string
	Keys       []string
	Actions    []Action
}

type Message struct {
	UserID   string
	Platform string
	Value    string
}

type Action struct {
	MessageType string
	Value       string
}
