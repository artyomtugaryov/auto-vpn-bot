package storage

type Customer struct {
	Username string
	Active   bool
}

type Message struct {
	Text string
}

type Paymant struct {
	Amount    float32
	Delivered bool
	Executed  bool
}
