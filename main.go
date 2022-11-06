package main

import (
	"fmt"

	"github.com/asaskevich/EventBus"
	"github.com/google/uuid"
	mod_foobar_services "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/foobar/services"
	mod_order_core_ents "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/order/core/entities"
	mod_order_core_events "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/order/core/events"
	mod_order_core_vos "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/order/core/value-objects"
	mod_shared "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/shared"
)

func main() {
	evtSubEvent1Handler := mod_foobar_services.NewOrderShippingAddressUpdatedHandler()
	//
	//
	//
	evtPub := mod_foobar_services.NewEventPublisher()
	evtPub.Subscribe(
		evtSubEvent1Handler,
		mod_order_core_events.NewEvent_OrderShippingAddressUpdated(mod_order_core_events.OrderShippingAddressUpdated_Attrs{}),
	)
	// evtPub.Subscribe(evtSubEvent1Handler, mod_order_core_events.NewEvent_OrderShippingAddressUpdated())
	// evtPub.Subscribe(evtSubEvent1Handler, mod_order_core_events.NewEvent_OrderShippingAddressUpdated())
	//
	//
	//

	bus := EventBus.New()
	for eventName, handlers := range evtPub.Events() {
		for _, handler := range handlers {
			bus.Subscribe(eventName, func(event mod_shared.DomainEvent) {
				handler.Notify(event)
			})
		}
	}

	// bus.Subscribe("main.calculator", calculator)
	// bus.Publish("main.calculator", 20, 49)
	// bus.Unsubscribe("main.calculator", calculator)
	// bus.Unsubscribe("event.order.shipping_address_updated", func() {
	// })

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
		bus.Publish(event.Name(), event)
	}
}

// func calculator(a int, b int) {
// 	fmt.Printf("%d\n", a+b)
// }
