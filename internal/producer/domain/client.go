package domain

import (
	"github.com/google/uuid"
	"github.com/rezamokaram/sample-ws/internal/common/domain"
)

type Client struct {
	ID      uuid.UUID
	Channel <-chan domain.Message
}
