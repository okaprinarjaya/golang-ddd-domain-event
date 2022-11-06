package mod_order_core_events

type OrderShippingAddressUpdated_Attrs struct {
	OrderId *string
}

type orderShippingAddressUpdated struct {
	orderId string
}

func NewEvent_OrderShippingAddressUpdated(attrs OrderShippingAddressUpdated_Attrs) OrderEvent {
	o := &orderShippingAddressUpdated{}
	if attrs.OrderId != nil {
		o.orderId = *attrs.OrderId
	}
	return o
}

func (evt *orderShippingAddressUpdated) Name() string {
	return "event.order.shipping_address_updated"
}

func (evt *orderShippingAddressUpdated) OrderId() string {
	return evt.orderId
}
