package mod_stock_core_ents_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	mod_shared "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/shared"
	ents "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/stock/core/entities"
)

// Domain / Business Entity Layer Testing
// Menambah jumlah stock item
// Item: "Groovy Soft Drink Root Beer 330mL" dengan onHandQty=10, reservedUnconfirmedQty=3, reservedConfirmedQty=3, availableQty=4

// Positive cases

type StockItemCommittingInTestSuite struct {
	suite.Suite
	StockItem *ents.StockItemEntity
	UserBy    string
}

func (suite *StockItemCommittingInTestSuite) SetupSuite() {
	suite.StockItem = ents.NewStockItemEntity()
	suite.StockItem.Id = "stock-item-001"
	suite.StockItem.CreatedAt = time.Now()
	suite.StockItem.CreatedBy = "SYSTEM"
	//
	suite.StockItem.ItemId = "item-001-groovy-soft-drink-root-beer-300-ml"
	suite.StockItem.OnHandQty = 10
	suite.StockItem.ReservedUnconfirmedQty = 3
	suite.StockItem.ReservedConfirmedQty = 3
	suite.StockItem.AvailableQty = 4
	suite.StockItem.DcId = "DC-001"
	suite.StockItem.StoreId = "STORE-001"
	suite.StockItem.OutOfStockThreshold = 3

	//
	suite.UserBy = "SYSTEM"
}

// Jumlah ditambahkan ke stock sebanyak "10" dengan nomer request: "STOCK-IN-001" dicommit
func (suite *StockItemCommittingInTestSuite) TestCommittingInStockItem_A_StockIn001_Amount10_Success() {
	suite.StockItem.CommitStockItemIN("STOCK-IN-001", 10, suite.UserBy)

	suite.Equal(3, suite.StockItem.ReservedUnconfirmedQty)
	suite.Equal(3, suite.StockItem.ReservedConfirmedQty)
	suite.Equal(20, suite.StockItem.OnHandQty)
	suite.Equal(14, suite.StockItem.AvailableQty)

	suite.NotNil(suite.StockItem.LatestCommittedStockItem)
	suite.IsType(ents.NewStockItemLogEntity(), suite.StockItem.LatestCommittedStockItem)
	suite.Equal(mod_shared.NEW, suite.StockItem.LatestCommittedStockItem.PersistenceStatus)
	suite.Equal(10, suite.StockItem.LatestCommittedStockItem.Qty)
	suite.Equal(20, suite.StockItem.LatestCommittedStockItem.OnHandQty)
}

func TestStockItemCommittingInTestSuite(t *testing.T) {
	suite.Run(t, new(StockItemCommittingInTestSuite))
}
