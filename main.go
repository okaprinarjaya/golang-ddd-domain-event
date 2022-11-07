package main

import (
	"fmt"

	"github.com/asaskevich/EventBus"
	"github.com/google/uuid"
	mod_foobar_subs "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/foobar/subscribers"
	mod_order_core_ents "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/order/core/entities"
	mod_order_core_vos "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/order/core/value-objects"
	mod_shared "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/shared"
)

func main() {
	// Application Bootstraping
	eventBus := EventBus.New()
	domainEventPublisher := mod_shared.NewEventPublisher()

	mod_foobar_subs.Subscriptions(eventBus, domainEventPublisher)

	for eventName, handlers := range domainEventPublisher.Events() {
		for _, handler := range handlers {
			eventBus.Subscribe(eventName, func(event mod_shared.DomainEvent) {
				handler.Notify(event)
			})
		}
	}

	//
	//
	//

	// Any Testing

	// Test Order core domain
	order := mod_order_core_ents.NewOrderEntity()
	order.DomainBaseEntity.Id = uuid.New().String()
	order.CreatedBy = "SYSTEM"
	order.CreatedByName = "SYSTEM"
	//
	order.CustomerId = "CUSTOMER-001"
	order.ItemId = "ITEM-001"
	order.InvoiceNumber = "INV-001-001-001"
	order.Status = mod_order_core_vos.OrderStatus_Draft
	order.TotalQuantity = 3
	order.TotalOrderValue = 500000
	order.ShippingCost = 75000
	order.IsCOD = false
	order.ShippingAddress = mod_order_core_vos.ShippingAddress{
		Street:  "Street",
		City:    "City",
		State:   "State",
		Country: "Indonesia",
		ZipCode: "7113",
	}

	event, err := order.UpdateShippingAddress(mod_order_core_vos.ShippingAddress{
		Street:  "Street updated",
		City:    "City updated",
		State:   "State updated",
		Country: "Indonesia updated",
		ZipCode: "7115",
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Trying to notify")
		eventBus.Publish(event.Name(), event)
	}
}
