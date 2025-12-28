package producer

import (
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	commonDomain "github.com/rezamokaram/sample-ws/internal/common/domain"
	"github.com/rezamokaram/sample-ws/internal/producer/domain"
	"github.com/rezamokaram/sample-ws/internal/producer/port"
)

type service struct {
	RWMutex *sync.RWMutex
	Clients map[uuid.UUID]client
}

type client struct {
	ID      uuid.UUID
	Mutex   *sync.Mutex
	Channel chan commonDomain.Message
}

func NewService() port.Service {
	svc := &service{
		RWMutex: new(sync.RWMutex),
		Clients: make(map[uuid.UUID]client),
	}

	// worker
	go svc.workerLoop()

	return svc
}

func (s *service) NewClient() (domain.Client, error) {
	client := client{
		ID:      uuid.New(),
		Mutex:   new(sync.Mutex),
		Channel: make(chan commonDomain.Message),
	}

	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()
	s.Clients[client.ID] = client

	return domain.Client{
		ID:      client.ID,
		Channel: client.Channel,
	}, nil
}

func (s *service) DeleteClient(client domain.Client) error {
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()
	delete(s.Clients, client.ID)

	return nil
}

func (s *service) workerLoop() {
	defer s.recoverWorkerFromPanic()

	data := uint(0)
	msg := commonDomain.Message{
		MessageType: 1,
		Data:        0,
	}
	for {
		msg.Data = data
		time.Sleep(20 * time.Millisecond)
		s.RWMutex.RLock()
		for _, client := range s.Clients {
			client.Channel <- msg
		}
		s.RWMutex.RUnlock()
		data++
	}
}

func (s *service) recoverWorkerFromPanic() {
	if r := recover(); r != nil {
		log.Printf("goroutine panic: %v\n", r)
		go s.workerLoop()
	}
	log.Printf("worker loop ends\n")
}
