package dispatcher_test

import (
	"testing"
	"github.com/jgimeno/eventdispatcher/dispatcher"
	"github.com/jgimeno/eventdispatcher/event"
)

func TestCreationOfDispatcher(t *testing.T) {
	d := dispatcher.New()

	eventName := "event.new"

	executedListener := false

	d.Subscribe(eventName, func(event event.Event) {
		executedListener = true
	})

	e := event.New(eventName)
	d.Publish(e)

	if !executedListener {
		t.Fatalf("The Dispatcher has not executed the listener.")
	}
}
