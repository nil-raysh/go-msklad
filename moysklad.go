package go_moysklad

import (
	"github.com/arcsub/go-moysklad/moysklad"
)

// ExtendedClient оборачивает moysklad.Client для добавления новых методов.
type ExtendedClient struct {
	*moysklad.Client
}

// NewExtendedMoySkladClient создает новый расширенный клиент.
func NewExtendedMoySkladClient() *ExtendedClient {
	// Инициализация оригинального клиента
	client := moysklad.NewClient()
	return &ExtendedClient{Client: client}
}

type EntityService struct {
	*moysklad.EntityService
}
