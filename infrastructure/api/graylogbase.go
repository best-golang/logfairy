package api

// GraylogBase defines the base functions
type GraylogBase interface {
	GetAuth(tokenName string) ([]string, error)
	HandleFailure(response []byte, status int) error
}
