package chat

type Chat interface {
	Create(name string, description string, userIds []string) (string, error)
}
