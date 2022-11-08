package mod_order_core_entities_tests

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	mod_order_core_ents "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/order/core/entities"
	mod_order_core_vos "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/order/core/value-objects"
)

// Positive cases

func TestUpdateShippingAddress_Success(t *testing.T) {
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

	assert.Nil(t, err)
	assert.NotNil(t, event)
	assert.Equal(t, "Street updated", order.ShippingAddress.Street)
	assert.Equal(t, mod_order_core_vos.OrderStatus_ShippingAddressConfirmed, int(order.Status))
}

// Negative cases

func TestUpdateShippingAddress_Failed_Because_Of_Status_Equal_Dispatched_Or_Above_Dispatched(t *testing.T) {
	order := mod_order_core_ents.NewOrderEntity()
	order.DomainBaseEntity.Id = uuid.New().String()
	order.CreatedBy = "SYSTEM"
	order.CreatedByName = "SYSTEM"
	//
	order.CustomerId = "CUSTOMER-001"
	order.ItemId = "ITEM-001"
	order.InvoiceNumber = "INV-001-001-001"
	order.Status = mod_order_core_vos.OrderStatus_Dispatched
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

	assert.NotNil(t, err)
	assert.Equal(t, "cannot update Shipping Address, order has already been dispatched", err.Error())
	assert.Nil(t, event)
	assert.Equal(t, "Street", order.ShippingAddress.Street)
	assert.Equal(t, mod_order_core_vos.OrderStatus_Dispatched, int(order.Status))
}
