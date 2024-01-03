package storage

type Storage interface {
	Initialize() error
	// Customer methods
	SaveCusromer(customer *Customer) error
	DisableCusromer(customer *Customer) error
	EnableCusromer(customer *Customer) error
}
