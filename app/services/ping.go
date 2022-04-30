package services

type Ping interface {
	Get() string
}

type ping struct {
}

func (p *ping) Get() string {
	return "Pong!!!!!!"
}

func NewPing() Ping {
	return &ping{}
}
