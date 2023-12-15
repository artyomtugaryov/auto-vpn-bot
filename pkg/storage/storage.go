package storage

type Storage interface {
	Initialize() error
	SaveCusromer(customer *Customer) error
}
