package session

type Repository interface {
	SetToken(id string, token string) error
	GetToken(id string) (string, error)
}
