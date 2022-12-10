package storage

type EventInterface interface {
	Create() bool
	Edit(int) int
	Delete(int) bool
}

type Event struct {
	ID    string
	Title string
	// TODO
}

type EventList *[]Event

func NewEvent(id, title string) EventInterface {
	return &Event{id, title}
}

func (e *Event) Edit(int) int {
	return 0
}

func (e *Event) Delete(int) bool {
	return true
}

func (e *Event) Create() bool {
	return true
}
