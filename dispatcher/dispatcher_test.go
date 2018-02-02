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

func TestWeCanSubscribeMultipleListenersToSameEvent(t *testing.T) {
	d := dispatcher.New()

	eventName := "event.new"

	var executedListener1 bool
	var executedListener2 bool

	var listener1 dispatcher.Listener = func(event event.Event, w *sync.WaitGroup) {
		executedListener1 = true
		w.Done()
	}

	var listener2 dispatcher.Listener = func(event event.Event, w *sync.WaitGroup) {
		executedListener2 = true
		w.Done()
	}

	d.Subscribe(eventName, listener1)
	d.Subscribe(eventName, listener2)

	e := event.New(eventName)
	d.Publish(e)
	d.Close()

	if !executedListener1 || !executedListener2 {
		t.Fatalf("The Dispatcher has not executed all the expected listeners.")
	}
}