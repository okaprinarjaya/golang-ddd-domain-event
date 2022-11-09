package mod_stock_core_ents_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	ents "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/stock/core/entities"
)

// Domain / Business Entity Layer Testing
// Membatalkan reservasi yang sudah dikonfirmasi
// Item: "Groovy Soft Drink Root Beer 330mL" dengan onHandQty=10, reservedUnconfirmedQty=1, reservedConfirmedQty=4, availableQty=5

// Positive cases

type StockItemCancelConfirmedReservationTestSuite struct {
	suite.Suite
	StockItem *ents.StockItemEntity
	UserBy    string
}

func (suite *StockItemCancelConfirmedReservationTestSuite) SetupSuite() {
	suite.StockItem = ents.NewStockItemEntity()
	suite.StockItem.Id = "stock-item-001"
	suite.StockItem.CreatedAt = time.Now()
	suite.StockItem.CreatedBy = "SYSTEM"
	//
	suite.StockItem.ItemId = "item-001-groovy-soft-drink-root-beer-300-ml"
	suite.StockItem.OnHandQty = 10
	suite.StockItem.ReservedUnconfirmedQty = 1
	suite.StockItem.ReservedConfirmedQty = 4
	suite.StockItem.AvailableQty = 5
	suite.StockItem.DcId = "DC-001"
	suite.StockItem.StoreId = "STORE-001"
	suite.StockItem.OutOfStockThreshold = 3

	//
	suite.UserBy = "SYSTEM"
}

// Membatalkan reservasi stock item yg terkonfirmasi sebanyak "2" dengan nomer: "ORDER-002"
func (suite *StockItemCancelConfirmedReservationTestSuite) TestCancellingConfirmedReservationOfStockItem_A_Order002_Amount1_Success() {
	suite.StockItem.CancelConfirmedReservationOfStockItem(2, suite.UserBy)

	suite.Nil(suite.StockItem.LatestCommittedStockItem)
	suite.Equal(1, suite.StockItem.ReservedUnconfirmedQty)
	suite.Equal(2, suite.StockItem.ReservedConfirmedQty)
	suite.Equal(10, suite.StockItem.OnHandQty)
	suite.Equal(7, suite.StockItem.AvailableQty)
}

// Membatalkan reservasi stock item yg terkonfirmasi sebanyak "2" dengan nomer: "ORDER-003"
func (suite *StockItemCancelConfirmedReservationTestSuite) TestCancellingConfirmedReservationOfStockItem_B_Order003_Amount2_Success() {
	suite.StockItem.CancelConfirmedReservationOfStockItem(2, suite.UserBy)

	suite.Nil(suite.StockItem.LatestCommittedStockItem)
	suite.Equal(1, suite.StockItem.ReservedUnconfirmedQty)
	suite.Equal(0, suite.StockItem.ReservedConfirmedQty)
	suite.Equal(10, suite.StockItem.OnHandQty)
	suite.Equal(9, suite.StockItem.AvailableQty)
}

func TestStockItemCancelConfirmedReservationTestSuite(t *testing.T) {
	suite.Run(t, new(StockItemCancelConfirmedReservationTestSuite))
}
