package main

import "github.com/google/uuid"

type Event interface {
	Name() string
}

type GeneralError string

func NewGeneralError(err error) Event {
	return GeneralError(err.Error())
}

func (ge GeneralError) Name() string {
	return "event.general.error"
}

// OrderEvent interface for describing Order relevant Domain Event
type OrderEvent interface {
	Event
	OrderID() uuid.UUID
}

type OrderDispatched struct {
	orderID uuid.UUID
}

func (odp OrderDispatched) Name() string {
	return "event.order.dispatched"
}

func (odp OrderDispatched) OrderID() uuid.UUID {
	return odp.orderID
}

// OrderDelivered actual event
type OrderDelivered struct {
	orderID uuid.UUID
}

func (odl OrderDelivered) Name() string {
	return "event.order.delivery.success"
}

func (odl OrderDelivered) OrderID() uuid.UUID {
	return odl.orderID
}

// OrderDeliveryFailed actual event
type OrderDeliveryFailed struct {
	orderID uuid.UUID
}

func (odlf OrderDeliveryFailed) Name() string {
	return "event.order.delivery.failed"
}

func (odlf OrderDeliveryFailed) OrderID() uuid.UUID {
	return odlf.orderID
}
