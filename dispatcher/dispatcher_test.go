package dispatcher_test

import (
	"testing"
	"github.com/jgimeno/eventdispatcher/dispatcher"
	"github.com/jgimeno/eventdispatcher/event"
	"sync"
)

func TestCreationOfDispatcher(t *testing.T) {
	d := dispatcher.New()

	eventName := "event.new"

	executedListener := false

	d.Subscribe(eventName, func(event event.Event, w *sync.WaitGroup) {
		executedListener = true
		w.Done()
	})

	e := event.New(eventName)
	d.Publish(e)
	d.End()

	if !executedListener {
		t.Fatalf("The Dispatcher has not executed the listener.")
	}
}
