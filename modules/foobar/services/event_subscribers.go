package mod_foobar_services

import (
	"fmt"

	mod_order_core_events "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/order/core/events"
	mod_shared "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/shared"
)

type DomainEventHandler interface {
	Notify(event mod_shared.DomainEvent)
}

type DomainEventPublisher struct {
	handlers map[string][]DomainEventHandler
}

func NewEventPublisher() *DomainEventPublisher {
	return &DomainEventPublisher{
		handlers: make(map[string][]DomainEventHandler),
	}
}

func (evt *DomainEventPublisher) Subscribe(handler DomainEventHandler, events ...mod_shared.DomainEvent) {
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

func (evt *DomainEventPublisher) Notify(event mod_shared.DomainEvent) {
	for _, handler := range evt.handlers[event.Name()] {
		handler.Notify(event)
	}
}

func (evt *DomainEventPublisher) Events() map[string][]DomainEventHandler {
	return evt.handlers
}

//
//
//

type OrderShippingAddressUpdatedHandler struct {
}

func NewOrderShippingAddressUpdatedHandler() DomainEventHandler {
	return DomainEventHandler(&OrderShippingAddressUpdatedHandler{})
}

func (h *OrderShippingAddressUpdatedHandler) Notify(event mod_shared.DomainEvent) {
	fmt.Printf("Notifyed!!!: %s\n", event.Name())
	fmt.Printf("Processing: %s\n", event.(mod_order_core_events.OrderEvent).OrderId())
}
