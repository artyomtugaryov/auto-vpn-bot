package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Process interface {
	process(event Event) error
}

type Type int

const (
	Unknow Type = iota
	Message
)

type Event struct {
	Type Type
}
