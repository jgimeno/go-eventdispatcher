package dispatcher

import (
	"sync"
	"github.com/jgimeno/go-eventdispatcher/event"
)

type Listener func(event event.Event, w *sync.WaitGroup)

type EventDispatcher interface {
	Publish(event event.Event)
	Subscribe(eventName string, listener Listener)
	Close()
}

func New() EventDispatcher {
	return &eventDispatcher{
		eventMap: make(map[string]Listener, 1),
	}
}

type eventDispatcher struct {
	eventMap map[string]Listener
	waitGroup sync.WaitGroup
}

func (e *eventDispatcher) Subscribe(eventName string, listener Listener) {
	e.eventMap[eventName] = listener
}

func (e *eventDispatcher) Publish(event event.Event) {
	e.waitGroup.Add(1)
	l := e.eventMap[event.GetName()]
	go l(event, &e.waitGroup)
}

func (e *eventDispatcher) Close() {
	e.waitGroup.Wait()
}
