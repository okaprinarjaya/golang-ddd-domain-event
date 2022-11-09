package mod_stock_core_ents_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	mod_shared "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/shared"
	ents "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/stock/core/entities"
)

// Domain / Business Entity Layer Testing
// Mengurangi jumlah stock item saat commit pesanan-pesanan yang sudah direservasi sebelumnya
// Item: "Groovy Soft Drink Root Beer 330mL" dengan onHandQty=10, reservedUnconfirmedQty=0, reservedConfirmedQty=5, availableQty=5

// Positive cases

type StockItemCommittingOutTestSuite struct {
	suite.Suite
	StockItem *ents.StockItemEntity
	UserBy    string
}

func (suite *StockItemCommittingOutTestSuite) SetupSuite() {
	suite.StockItem = ents.NewStockItemEntity()
	suite.StockItem.Id = "stock-item-001"
	suite.StockItem.CreatedAt = time.Now()
	suite.StockItem.CreatedBy = "SYSTEM"
	//
	suite.StockItem.ItemId = "item-001-groovy-soft-drink-root-beer-300-ml"
	suite.StockItem.OnHandQty = 10
	suite.StockItem.ReservedUnconfirmedQty = 0
	suite.StockItem.ReservedConfirmedQty = 5
	suite.StockItem.AvailableQty = 5
	suite.StockItem.DcId = "DC-001"
	suite.StockItem.StoreId = "STORE-001"
	suite.StockItem.OutOfStockThreshold = 3

	//
	suite.UserBy = "SYSTEM"
}

// Item dikeluarkan dari stock sebanyak "1" saat pesanan dengan nomer: "ORDER-001" dicommit
func (suite *StockItemCommittingOutTestSuite) TestCommittingOutStockItem_A_Order001_Amount1_Success() {
	suite.StockItem.CommitStockItemOUT("ORDER-001", 1, suite.UserBy)

	suite.Equal(0, suite.StockItem.ReservedUnconfirmedQty)
	suite.Equal(4, suite.StockItem.ReservedConfirmedQty)
	suite.Equal(9, suite.StockItem.OnHandQty)
	suite.Equal(5, suite.StockItem.AvailableQty)

	suite.NotNil(suite.StockItem.LatestCommittedStockItem)
	suite.IsType(ents.NewStockItemLogEntity(), suite.StockItem.LatestCommittedStockItem)
	suite.Equal(mod_shared.NEW, suite.StockItem.LatestCommittedStockItem.PersistenceStatus)
	suite.Equal(1, suite.StockItem.LatestCommittedStockItem.Qty)
	suite.Equal(9, suite.StockItem.LatestCommittedStockItem.OnHandQty)
}

// Item dikeluarkan dari stock sebanyak "2" saat pesanan dengan nomer: "ORDER-002" dicommit
func (suite *StockItemCommittingOutTestSuite) TestCommittingOutStockItem_B_Order002_Amount2_Success() {
	suite.StockItem.CommitStockItemOUT("ORDER-002", 2, suite.UserBy)

	suite.Equal(0, suite.StockItem.ReservedUnconfirmedQty)
	suite.Equal(2, suite.StockItem.ReservedConfirmedQty)
	suite.Equal(7, suite.StockItem.OnHandQty)
	suite.Equal(5, suite.StockItem.AvailableQty)

	suite.NotNil(suite.StockItem.LatestCommittedStockItem)
	suite.IsType(ents.NewStockItemLogEntity(), suite.StockItem.LatestCommittedStockItem)
	suite.Equal(mod_shared.NEW, suite.StockItem.LatestCommittedStockItem.PersistenceStatus)
	suite.Equal(2, suite.StockItem.LatestCommittedStockItem.Qty)
	suite.Equal(7, suite.StockItem.LatestCommittedStockItem.OnHandQty)
}

// Item dikeluarkan dari stock sebanyak "2" saat pesanan dengan nomer: "ORDER-003" dicommit
func (suite *StockItemCommittingOutTestSuite) TestCommittingOutStockItem_C_Order003_Amount2_Success() {
	suite.StockItem.CommitStockItemOUT("ORDER-003", 2, suite.UserBy)

	suite.Equal(0, suite.StockItem.ReservedUnconfirmedQty)
	suite.Equal(0, suite.StockItem.ReservedConfirmedQty)
	suite.Equal(5, suite.StockItem.OnHandQty)
	suite.Equal(5, suite.StockItem.AvailableQty)

	suite.NotNil(suite.StockItem.LatestCommittedStockItem)
	suite.IsType(ents.NewStockItemLogEntity(), suite.StockItem.LatestCommittedStockItem)
	suite.Equal(mod_shared.NEW, suite.StockItem.LatestCommittedStockItem.PersistenceStatus)
	suite.Equal(2, suite.StockItem.LatestCommittedStockItem.Qty)
	suite.Equal(5, suite.StockItem.LatestCommittedStockItem.OnHandQty)
}

func TestStockItemCommittingOutTestSuite(t *testing.T) {
	suite.Run(t, new(StockItemCommittingOutTestSuite))
}
