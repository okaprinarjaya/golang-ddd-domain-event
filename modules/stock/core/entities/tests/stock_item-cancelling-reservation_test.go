package mod_stock_core_ents_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	ents "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/stock/core/entities"
)

// Domain / Business Entity Layer Testing
// Membatalkan reservasi
// Item: \"Groovy Soft Drink Root Beer 330mL\" dengan onHandQty=10, reservedUnconfirmedQty=5, reservedConfirmedQty=0, availableQty=5

// Positive cases

type StockItemCancellingReservationTestSuite struct {
	suite.Suite
	StockItem *ents.StockItemEntity
	UserBy    string
}

func (suite *StockItemCancellingReservationTestSuite) SetupSuite() {
	suite.StockItem = ents.NewStockItemEntity()
	suite.StockItem.Id = "stock-item-001"
	suite.StockItem.CreatedAt = time.Now()
	suite.StockItem.CreatedBy = "SYSTEM"
	//
	suite.StockItem.ItemId = "item-001-groovy-soft-drink-root-beer-300-ml"
	suite.StockItem.OnHandQty = 10
	suite.StockItem.ReservedUnconfirmedQty = 5
	suite.StockItem.ReservedConfirmedQty = 0
	suite.StockItem.AvailableQty = 5
	suite.StockItem.DcId = "DC-001"
	suite.StockItem.StoreId = "STORE-001"
	suite.StockItem.OutOfStockThreshold = 3

	//
	suite.UserBy = "SYSTEM"
}

// Membatalkan reservasi stock item sebanyak "1" dengan nomer: "ORDER-001"
func (suite *StockItemCancellingReservationTestSuite) TestCancellingReservationOfStockItem_A_Order001_Amount1_Success() {
	suite.StockItem.CancelReservationOfStockItem(1, suite.UserBy)

	suite.Nil(suite.StockItem.LatestCommittedStockItem)
	suite.Equal(4, suite.StockItem.ReservedUnconfirmedQty)
	suite.Equal(0, suite.StockItem.ReservedConfirmedQty)
	suite.Equal(10, suite.StockItem.OnHandQty)
	suite.Equal(6, suite.StockItem.AvailableQty)
}

// Membatalkan reservasi stock item sebanyak "2" dengan nomer: "ORDER-003"
func (suite *StockItemCancellingReservationTestSuite) TestCancellingReservationOfStockItem_B_Order003_Amount2_Success() {
	suite.StockItem.CancelReservationOfStockItem(2, suite.UserBy)

	suite.Nil(suite.StockItem.LatestCommittedStockItem)
	suite.Equal(2, suite.StockItem.ReservedUnconfirmedQty)
	suite.Equal(0, suite.StockItem.ReservedConfirmedQty)
	suite.Equal(10, suite.StockItem.OnHandQty)
	suite.Equal(8, suite.StockItem.AvailableQty)
}

func TestStockItemCancellingReservationTestSuite(t *testing.T) {
	suite.Run(t, new(StockItemCancellingReservationTestSuite))
}
