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
		eventMap: make(map[string][]Listener, 1),
	}
}

type eventDispatcher struct {
	eventMap map[string][]Listener
	waitGroup sync.WaitGroup
}

func (e *eventDispatcher) Subscribe(eventName string, listener Listener) {
	e.eventMap[eventName] = append(e.eventMap[eventName], listener)
}

func (e *eventDispatcher) Publish(event event.Event) {
	for _, l := range e.eventMap[event.GetName()] {
		e.waitGroup.Add(1)
		go l(event, &e.waitGroup)
	}
}

func (e *eventDispatcher) Close() {
	e.waitGroup.Wait()
}
