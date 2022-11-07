package mod_shared

type DomainEvent interface {
	Name() string
}

type DomainEventHandler interface {
	Notify(event DomainEvent)
}

type DomainEventPublisher struct {
	handlers map[string][]DomainEventHandler
}

func NewEventPublisher() *DomainEventPublisher {
	return &DomainEventPublisher{
		handlers: make(map[string][]DomainEventHandler),
	}
}

func (evt *DomainEventPublisher) Subscribe(handler DomainEventHandler, events ...DomainEvent) {
	for _, event := range events {
		handlers := evt.handlers[event.Name()]
		if handlers != nil {
			handlers = append(handlers, handler)
			evt.handlers[event.Name()] = handlers
		} else {
			handlers := []DomainEventHandler{handler}
			evt.handlers = map[string][]DomainEventHandler{event.Name(): handlers}
		}
	}
}

func (evt *DomainEventPublisher) Notify(event DomainEvent) {
	for _, handler := range evt.handlers[event.Name()] {
		handler.Notify(event)
	}
}

func (evt *DomainEventPublisher) Events() map[string][]DomainEventHandler {
	return evt.handlers
}
