package mod_stock_core_ents

import mod_shared "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/shared"

type StockItemEntity struct {
	mod_shared.DomainBaseEntity
	ItemId                   string
	DcId                     string
	StoreId                  string
	AvailableQty             int
	OnHandQty                int
	ReservedUnconfirmedQty   int
	ReservedConfirmedQty     int
	OutOfStockThreshold      int
	LatestCommittedStockItem *StockItemLogEntity
}

func NewStockItemEntity() *StockItemEntity {
	return &StockItemEntity{}
}

func (ent *StockItemEntity) CreateStockItem() {
	ent.PersistenceStatus = mod_shared.NEW
	ent.createStockItemLog("INIT", "STOCK_IN", ent.OnHandQty)
}

func (ent *StockItemEntity) ReserveStockItem(qty int, userBy string) {
	ent.ReservedUnconfirmedQty = ent.ReservedUnconfirmedQty + qty
	ent.AvailableQty = ent.OnHandQty - ent.ReservedUnconfirmedQty - ent.ReservedConfirmedQty
	ent.LatestCommittedStockItem = nil
	ent.UpdatedBy = userBy
	ent.PersistenceStatus = mod_shared.MODIFIED
}

func (ent *StockItemEntity) ConfirmReservationOfStockItem(qty int, userBy string) {
	ent.ReservedUnconfirmedQty = ent.ReservedUnconfirmedQty - qty
	ent.ReservedConfirmedQty = ent.ReservedConfirmedQty + qty
	ent.AvailableQty = ent.OnHandQty - ent.ReservedUnconfirmedQty - ent.ReservedConfirmedQty
	ent.LatestCommittedStockItem = nil
	ent.UpdatedBy = userBy
	ent.PersistenceStatus = mod_shared.MODIFIED
}

func (ent *StockItemEntity) CancelReservationOfStockItem(qty int, userBy string) {
	ent.ReservedUnconfirmedQty = ent.ReservedUnconfirmedQty - qty
	ent.AvailableQty = ent.OnHandQty - ent.ReservedUnconfirmedQty - ent.ReservedConfirmedQty
	ent.LatestCommittedStockItem = nil
	ent.UpdatedBy = userBy
	ent.PersistenceStatus = mod_shared.MODIFIED
}

func (ent *StockItemEntity) CancelConfirmedReservationOfStockItem(qty int, userBy string) {
	ent.ReservedConfirmedQty = ent.ReservedConfirmedQty - qty
	ent.AvailableQty = ent.OnHandQty - ent.ReservedUnconfirmedQty - ent.ReservedConfirmedQty
	ent.LatestCommittedStockItem = nil
	ent.UpdatedBy = userBy
	ent.PersistenceStatus = mod_shared.MODIFIED
}

func (ent *StockItemEntity) CommitStockItemOUT(documentId string, qty int, userBy string) {
	ent.ReservedConfirmedQty = ent.ReservedConfirmedQty - qty
	ent.OnHandQty = ent.OnHandQty - qty
	ent.AvailableQty = ent.OnHandQty - ent.ReservedUnconfirmedQty - ent.ReservedConfirmedQty
	ent.UpdatedBy = userBy
	ent.PersistenceStatus = mod_shared.MODIFIED

	ent.createStockItemLog(documentId, "STOCK_OUT", qty)
}

func (ent *StockItemEntity) CommitStockItemIN(documentId string, qty int, userBy string) {
	ent.OnHandQty = ent.OnHandQty + qty
	ent.AvailableQty = ent.OnHandQty - ent.ReservedUnconfirmedQty - ent.ReservedConfirmedQty
	ent.UpdatedBy = userBy
	ent.PersistenceStatus = mod_shared.MODIFIED

	ent.createStockItemLog(documentId, "STOCK_IN", qty)
}

func (ent *StockItemEntity) createStockItemLog(documentId string, trxType string, qty int) {
	stockItemLogBizEnt := NewStockItemLogEntity()
	stockItemLogBizEnt.CreatedBy = "SYSTEM"
	stockItemLogBizEnt.StockItemId = ent.Id
	stockItemLogBizEnt.DocumentId = documentId
	stockItemLogBizEnt.ItemId = ent.ItemId
	stockItemLogBizEnt.DcId = ent.DcId
	stockItemLogBizEnt.StoreId = ent.StoreId
	stockItemLogBizEnt.TransactionType = trxType
	stockItemLogBizEnt.Qty = qty
	stockItemLogBizEnt.OnHandQty = ent.OnHandQty

	stockItemLogBizEnt.PersistenceStatus = mod_shared.NEW

	ent.LatestCommittedStockItem = stockItemLogBizEnt
}
