package mod_order_core_ents

import (
	"fmt"

	order_core_evts "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/order/core/events"
	order_core_vos "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/order/core/value-objects"
	mod_shared "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/shared"
)

type OrderEntity struct {
	mod_shared.DomainBaseEntity
	CustomerId      string
	ItemId          string
	InvoiceNumber   string
	Status          order_core_vos.OrderStatus
	TotalQuantity   int
	TotalOrderValue int
	ShippingCost    int
	IsCOD           bool
	ShippingAddress order_core_vos.ShippingAddress
}

func NewOrderEntity() *OrderEntity {
	return &OrderEntity{}
}

func (ent *OrderEntity) UpdateShippingAddress(shippingAddr order_core_vos.ShippingAddress) (mod_shared.DomainEvent, error) {
	if ent.Status >= order_core_vos.OrderStatus_Dispatched {
		return nil, fmt.Errorf("cannot update Shipping Address, order has already been dispatched")
	}
	if ent.Status < order_core_vos.OrderStatus_ShippingAddressConfirmed {
		ent.Status = order_core_vos.OrderStatus_ShippingAddressConfirmed
	}
	ent.ShippingAddress = shippingAddr
	ent.PersistenceStatus = mod_shared.MODIFIED

	event := order_core_evts.NewEvent_OrderShippingAddressUpdated(order_core_evts.OrderShippingAddressUpdated_Attrs{
		OrderId: &ent.Id,
	})

	return event, nil
}
