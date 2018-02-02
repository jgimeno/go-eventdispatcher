package event

type Event interface {
	GetName() string
}

type event struct {
	name string
}

func (e event) GetName() string {
	return e.name
}

func New(name string) Event {
	return event{name}
}
