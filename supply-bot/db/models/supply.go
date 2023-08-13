package models

type SupplyStatus uint
type OrderType uint

const (
	SupplyStatusInit           SupplyStatus = 1
	SupplyStatusPlaced         SupplyStatus = 2
	SupplyStatusFulfilled      SupplyStatus = 3
	SupplyStatusExpired        SupplyStatus = 4
	SupplyStatusrequestFulfill SupplyStatus = 5
)

type Supply struct {
	Supplier       string `gorm:"primaryKey;" json:"supplier"`
	Token          string `json:"token"`
	TotalAmount    uint64 `json:"total_amount"`
	SupplyAmount   uint64 `json:"supply_amount"`
	MinAmount      uint64 `json:"min_amount"`
	SplitTime      uint64 `json:"split_time"`
	LastSupplyTime uint64 `json:"last_supply_time"`
	Reciever       string `json:"reciever"`
	TxSubmitted    bool   `json:"tx_submitted"`
}
