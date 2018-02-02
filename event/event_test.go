package event_test

import (
	"github.com/jgimeno/eventdispatcher/event"
	"testing"
)

func TestWeCanCreateAnEvent(t *testing.T) {
	e := event.New("Name")

	if e.GetName() != "Name" {
		t.Fatalf("Error getting the name of the event.")
	}
}
