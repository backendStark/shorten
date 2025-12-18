package stat

import (
	"log"
	"shorten/pkg/event"
)

type StatServiceDeps struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

type StatService struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

func NewAuthService(deps StatServiceDeps) *StatService {
	return &StatService{
		EventBus:       deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

func (s *StatService) AddClick() {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Fatalln("Bad EventLinkVisited type", msg.Data)
				continue
			}
			s.StatRepository.AddClick(id)
		}
	}
}
