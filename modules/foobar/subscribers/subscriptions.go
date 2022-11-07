package mod_foobar_subscribers

import (
	"github.com/asaskevich/EventBus"
	mod_order_core_events "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/order/core/events"
	mod_shared "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/shared"
)

func Subscriptions(eventBus EventBus.Bus, domainEventPublisher *mod_shared.DomainEventPublisher) {
	event1 := mod_order_core_events.NewEvent_OrderShippingAddressUpdated(mod_order_core_events.OrderShippingAddressUpdated_Attrs{})
	handler1 := NewOrderShippingAddressUpdatedHandler()

	event2 := mod_order_core_events.NewEvent_OrderShippingAddressUpdated(mod_order_core_events.OrderShippingAddressUpdated_Attrs{})
	handler2 := NewOrderShippingAddressUpdatedHandler()

	event3 := mod_order_core_events.NewEvent_OrderShippingAddressUpdated(mod_order_core_events.OrderShippingAddressUpdated_Attrs{})
	handler3 := NewOrderShippingAddressUpdatedHandler()

	domainEventPublisher.Subscribe(handler1, event1)
	domainEventPublisher.Subscribe(handler2, event2)
	domainEventPublisher.Subscribe(handler3, event3)
}
