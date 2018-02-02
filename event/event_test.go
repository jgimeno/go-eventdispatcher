package event_test

import (
	"testing"
	"github.com/jgimeno/go-eventdispatcher/event"
)

func TestWeCanCreateAnEvent(t *testing.T) {
	e := event.New("Name")

	if e.GetName() != "Name" {
		t.Fatalf("Error getting the name of the event.")
	}
}
