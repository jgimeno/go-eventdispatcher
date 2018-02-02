package dispatcher

import "github.com/jgimeno/eventdispatcher/event"

type Listener func(event event.Event)

type EventDispatcher interface {
	Publish(event event.Event)
	Subscribe(eventName string, listener Listener)
}

func New() EventDispatcher {
	return &eventDispatcher{
		eventMap: make(map[string]Listener, 1),
	}
}

type eventDispatcher struct {
	eventMap map[string]Listener
}

func (e *eventDispatcher) Subscribe(eventName string, listener Listener) {
	panic("implement me")
}

func (e *eventDispatcher) Publish(event event.Event) {
	panic("implement me")
}

