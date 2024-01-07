package storage

import "github.com/artyomtugaryov/vpnbot/pkg/entities"

type Storage interface {
	Initialize() error
	Close() error
	// Customer methods
	SaveCusromer(customer *entities.Customer) error
	DisableCusromer(customer *entities.Customer) error
	EnableCusromer(customer *entities.Customer) error
}
