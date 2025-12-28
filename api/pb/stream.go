package pb

import commonDomain "github.com/rezamokaram/sample-ws/internal/common/domain"

type Stream struct {
	Channel <-chan commonDomain.Message
}
