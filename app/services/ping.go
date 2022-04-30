package services

type PingService interface {
	Get() string
}

type pingService struct {
}

func (p *pingService) Get() string {
	return "Pong!!!!!!"
}

func NewPingService() PingService {
	return &pingService{}
}
