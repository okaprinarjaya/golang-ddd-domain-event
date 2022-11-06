package mod_order_core_vos

type OrderStatus int

const (
	OrderStatus_Draft                    = 0
	OrderStatus_ShippingAddressConfirmed = 100
	OrderStatus_PaymentMethodConfirmed   = 120
	OrderStatus_Submitted                = 150
	OrderStatus_Paid                     = 200
	OrderStatus_Dispatched               = 300
	OrderStatus_Delivered                = 400
	OrderStatus_Cancelled                = 900
)
