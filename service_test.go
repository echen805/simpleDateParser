package simpleDateParser

import (
	"context"
	"testing"
	"time"
)

func TestStatus(t *testing.T) {
	srv, ctx := setup()
	expected := "ok"

	s, err := srv.Status(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// Testing status
	if s != "ok" {
		t.Errorf("Status is invalid. Expected: %s, Actual: %s", expected, s)
	}
}

func TestGet(t *testing.T){
	srv, ctx := setup()
	currentDate, err := srv.Get(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	currentTime := time.Now()
	expected := currentTime.Format(layoutISO)

	if expected != currentDate {
		t.Errorf("Retrieved the wrong date. Expected: %s, Actual: %s", expected, currentDate)
	}
}

func TestValidate(t *testing.T) {
	srv, ctx := setup()
	b, err := srv.Validate(ctx, "20191231")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// testing an invalid date
	b, err = srv.Validate(ctx, "20193131")
	if b {
		t.Errorf("date should be invalid")
	}
}

func setup() (srv Service, ctx context.Context) {
	return NewService(), context.Background()
}
