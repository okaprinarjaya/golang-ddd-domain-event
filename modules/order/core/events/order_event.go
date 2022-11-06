package mod_order_core_events

import mod_shared "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/shared"

type OrderEvent interface {
	mod_shared.DomainEvent
	OrderId() string
}
