package mod_stock_core_ents_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	ents "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/stock/core/entities"
)

// Domain / Business Entity Layer Testing
// Membuat stock item baru
// Item: "Groovy Soft Drink Root Beer 330mL" dengan onHandQty=10, reservedUnconfirmedQty=0, reservedConfirmedQty=0, availableQty=10

// Positive cases

// Membuat satu stock item baru untuk item: "Groovy Soft Drink Root Beer 330mL" dengan inisial jumlah stock 10
func TestReserveStockItem_A_Order001_Amount1_Success(t *testing.T) {
	stockItem := ents.NewStockItemEntity()
	stockItem.Id = "stock-item-001"
	stockItem.CreatedAt = time.Now()
	stockItem.CreatedBy = "SYSTEM"
	//
	stockItem.ItemId = "item-001-groovy-soft-drink-root-beer-300-ml"
	stockItem.OnHandQty = 10
	stockItem.ReservedUnconfirmedQty = 0
	stockItem.ReservedConfirmedQty = 0
	stockItem.AvailableQty = 10
	stockItem.DcId = "DC-001"
	stockItem.StoreId = "STORE-001"
	stockItem.OutOfStockThreshold = 3

	stockItem.CreateStockItem()

	assert.NotNil(t, stockItem.LatestCommittedStockItem)
	assert.Equal(t, 0, stockItem.ReservedUnconfirmedQty)
	assert.Equal(t, 0, stockItem.ReservedConfirmedQty)
	assert.Equal(t, 10, stockItem.OnHandQty)
	assert.Equal(t, 10, stockItem.AvailableQty)
}
