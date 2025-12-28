package port

import "github.com/rezamokaram/sample-ws/internal/producer/domain"

type Service interface {
	NewClient() (domain.Client, error)
}
