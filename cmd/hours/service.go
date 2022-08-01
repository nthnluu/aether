package hours

import (
	proto "aether/pb/out"
	"errors"
	"strings"
	"time"
)

type Service interface {
	CreateQueue(courseId string, title string, description string, location string, endTime time.Time) (string, error)
	GetQueue(queueId string) (string, error)
	UpdateQueue(title string, description string, location string, endTime time.Time) (string, error)
	DeleteQueue(queueId string) (string, error)
	CreateTicket(courseId string, title string, description string, location string, endTime time.Time) (string, error)
	UpdateTicket(title string, description string, location string, endTime time.Time) (string, error)
	DeleteTicket(queueId string) (string, error)
}

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	var md proto.CreationMetadata
	md.GetCreatedBy()
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {

	return len(s)
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("empty string")
