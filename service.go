package simpleDateParser

import (
	"context"
	"time"
)

// Service provides some "date capabilities" to your application
type Service interface {
	Status(ctx context.Context) (string, error)
	Get(ctx context.Context) (string, error)
	Validate(ctx context.Context, date string) (bool, error)
}

type dateService struct{}


// NewService makes a new Service
func NewService() Service {
	return dateService{}
}

// Status returns bool if service is available
/* Tip:
(dateservice) is a receiver, so it would be called by a dateservice struct.
Warning: When called, the method is applied to a copy of the struct calling it.
*/
func (dateService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

// Get will return today's date
func (dateService) Get(ctx context.Context) (string, error) {
	now := time.Now()
	return now.Format(layoutISO), nil
}

const (
	layoutISO = "20060102"
	layoutUS  = "January 2, 2006"
)
// Validate will check if the input is the current date
func (dateService) Validate(ctx context.Context, date string) (bool, error) {
	currentTime := time.Now()
	currentDate := currentTime.Format(layoutISO)
	parsedDate, err := time.Parse(layoutISO, date)
	if err != nil {
		return false, err
	}

	if currentDate != parsedDate.Format(layoutISO) {
		return false, nil
	}
	return true, nil
}
