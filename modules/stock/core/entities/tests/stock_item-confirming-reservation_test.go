package mod_stock_core_ents_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	ents "gitlab.com/okaprinarjaya.wartek/ddd-domain-event/modules/stock/core/entities"
)

// Domain / Business Entity Layer Testing
// Mengkonfirmasi reservasi stock item
// Item: "Groovy Soft Drink Root Beer 330mL" dengan onHandQty=10, reservedUnconfirmedQty=0, reservedConfirmedQty=0, availableQty=10

// Positive cases

type StockItemConfirmingReservationTestSuite struct {
	suite.Suite
	StockItem *ents.StockItemEntity
	UserBy    string
}

func (suite *StockItemConfirmingReservationTestSuite) SetupSuite() {
	suite.StockItem = ents.NewStockItemEntity()
	suite.StockItem.Id = "stock-item-001"
	suite.StockItem.CreatedAt = time.Now()
	suite.StockItem.CreatedBy = "SYSTEM"
	//
	suite.StockItem.ItemId = "item-001-groovy-soft-drink-root-beer-300-ml"
	suite.StockItem.OnHandQty = 10
	suite.StockItem.ReservedUnconfirmedQty = 0
	suite.StockItem.ReservedConfirmedQty = 0
	suite.StockItem.AvailableQty = 10
	suite.StockItem.DcId = "DC-001"
	suite.StockItem.StoreId = "STORE-001"
	suite.StockItem.OutOfStockThreshold = 3

	//
	suite.UserBy = "SYSTEM"
}

// Mereservasi lalu langsung mengkonfirmasi

// Reservasi stock item sebanyak "1" lalu langsung mengkonfirmasi dengan nomer: "ORDER-001"
func (suite *StockItemConfirmingReservationTestSuite) TestConfirmReservationOfStockItem_A_Order001_Amount1_Success() {
	suite.StockItem.ReserveStockItem(1, suite.UserBy)

	suite.Nil(suite.StockItem.LatestCommittedStockItem)
	suite.Equal(1, suite.StockItem.ReservedUnconfirmedQty)
	suite.Equal(0, suite.StockItem.ReservedConfirmedQty)
	suite.Equal(10, suite.StockItem.OnHandQty)
	suite.Equal(9, suite.StockItem.AvailableQty)

	suite.StockItem.ConfirmReservationOfStockItem(1, suite.UserBy)

	suite.Nil(suite.StockItem.LatestCommittedStockItem)
	suite.Equal(0, suite.StockItem.ReservedUnconfirmedQty)
	suite.Equal(1, suite.StockItem.ReservedConfirmedQty)
	suite.Equal(10, suite.StockItem.OnHandQty)
	suite.Equal(9, suite.StockItem.AvailableQty)
}

// Mereservasi terlebih dahulu lalu disusul mengkonfirmasi kemudian

// Reservasi stock item sebanyak "2" dengan nomer: "ORDER-002"
func (suite *StockItemConfirmingReservationTestSuite) TestConfirmReservationOfStockItem_B_Reserve_Success() {
	suite.StockItem.ReserveStockItem(2, suite.UserBy)

	suite.Nil(suite.StockItem.LatestCommittedStockItem)
	suite.Equal(2, suite.StockItem.ReservedUnconfirmedQty)
	suite.Equal(1, suite.StockItem.ReservedConfirmedQty)
	suite.Equal(10, suite.StockItem.OnHandQty)
	suite.Equal(7, suite.StockItem.AvailableQty)
}

// Reservasi stock item sebanyak "2" dengan nomer: "ORDER-003"
func (suite *StockItemConfirmingReservationTestSuite) TestConfirmReservationOfStockItem_C_Reserve_Success() {
	suite.StockItem.ReserveStockItem(2, suite.UserBy)

	suite.Nil(suite.StockItem.LatestCommittedStockItem)
	suite.Equal(4, suite.StockItem.ReservedUnconfirmedQty)
	suite.Equal(1, suite.StockItem.ReservedConfirmedQty)
	suite.Equal(10, suite.StockItem.OnHandQty)
	suite.Equal(5, suite.StockItem.AvailableQty)
}

// Mengkonfirmasi reservasi sebanyak "2" dari order: "ORDER-002"
func (suite *StockItemConfirmingReservationTestSuite) TestConfirmReservationOfStockItem_D_Confirm_Success() {
	suite.StockItem.ConfirmReservationOfStockItem(2, suite.UserBy)

	suite.Nil(suite.StockItem.LatestCommittedStockItem)
	suite.Equal(2, suite.StockItem.ReservedUnconfirmedQty)
	suite.Equal(3, suite.StockItem.ReservedConfirmedQty)
	suite.Equal(10, suite.StockItem.OnHandQty)
	suite.Equal(5, suite.StockItem.AvailableQty)
}

// Mengkonfirmasi reservasi sebanyak "2" dari order: "ORDER-003"
func (suite *StockItemConfirmingReservationTestSuite) TestConfirmReservationOfStockItem_E_Confirm_Success() {
	suite.StockItem.ConfirmReservationOfStockItem(2, suite.UserBy)

	suite.Nil(suite.StockItem.LatestCommittedStockItem)
	suite.Equal(0, suite.StockItem.ReservedUnconfirmedQty)
	suite.Equal(5, suite.StockItem.ReservedConfirmedQty)
	suite.Equal(10, suite.StockItem.OnHandQty)
	suite.Equal(5, suite.StockItem.AvailableQty)
}

func TestStockItemConfirmingReservationTestSuite(t *testing.T) {
	suite.Run(t, new(StockItemConfirmingReservationTestSuite))
}
