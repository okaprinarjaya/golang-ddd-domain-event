package mod_order_core_ents

import (
	"fmt"

	mod_order_core_vos "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/order/core/value-objects"
	mod_shared "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/shared"
)

type OrderBizEntity struct {
	mod_shared.BaseEntity
	CustomerId      string
	ItemId          string
	InvoiceNumber   string
	Status          mod_order_core_vos.OrderStatus
	TotalQuantity   int
	TotalOrderValue int
	ShippingCost    int
	IsCOD           bool
	ShippingAddress mod_order_core_vos.ShippingAddress
}

func NewOrderBizEntity() *OrderBizEntity {
	return &OrderBizEntity{}
}

func (ent *OrderBizEntity) UpdateShippingAddress(shippingAddr mod_order_core_vos.ShippingAddress) error {
	if ent.Status >= mod_order_core_vos.OrderStatus_Dispatched {
		return fmt.Errorf("cannot update Shipping Address, order has already been dispatched")
	}
	if ent.Status < mod_order_core_vos.OrderStatus_ShippingAddressConfirmed {
		ent.Status = mod_order_core_vos.OrderStatus_ShippingAddressConfirmed
	}
	ent.ShippingAddress = shippingAddr

	return nil
}
