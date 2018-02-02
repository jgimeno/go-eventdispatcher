package dispatcher_test

import (
	"testing"
	"sync"
	"github.com/jgimeno/go-eventdispatcher/dispatcher"
	"github.com/jgimeno/go-eventdispatcher/event"
)

func TestItWaitsUntilItFinishesToCloseTheDispatcher(t *testing.T) {
	d := dispatcher.New()

	eventName := "event.new"

	executedListener := false

	d.Subscribe(eventName, func(event event.Event, w *sync.WaitGroup) {
		executedListener = true
		w.Done()
	})

	e := event.New(eventName)
	d.Publish(e)
	d.Close()

	if !executedListener {
		t.Fatalf("The Dispatcher has not executed the listener.")
	}
}
