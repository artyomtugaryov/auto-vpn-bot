package entities

type Customer struct {
	ID       uint8
	Username string
	Active   bool
}

type Message struct {
	ID   uint8
	Text string
}

type Paymant struct {
	ID        uint8
	Amount    float32
	Delivered bool
	Executed  bool
}
