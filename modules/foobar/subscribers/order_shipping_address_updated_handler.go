package mod_foobar_subscribers

import (
	"fmt"

	mod_order_core_events "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/order/core/events"
	mod_shared "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/shared"
)

type OrderShippingAddressUpdatedHandler struct {
}

func NewOrderShippingAddressUpdatedHandler() mod_shared.DomainEventHandler {
	return &OrderShippingAddressUpdatedHandler{}
}

func (h *OrderShippingAddressUpdatedHandler) Notify(event mod_shared.DomainEvent) {
	fmt.Printf("Notifyed!!!: %s\n", event.Name())
	fmt.Printf("Processing: %s\n", event.(mod_order_core_events.OrderEvent).OrderId())
	fmt.Println()
}
