package mod_stock_core_ents

import mod_shared "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/shared"

type StockItemLogEntity struct {
	mod_shared.DomainBaseEntity
	StockItemId     string
	DocumentId      string
	ItemId          string
	DcId            string
	StoreId         string
	TransactionType string
	Qty             int
	OnHandQty       int
}

func NewStockItemLogEntity() *StockItemLogEntity {
	return &StockItemLogEntity{}
}
