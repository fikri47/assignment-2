package entity

type Item struct {
	ItemID      uint   `gorm:"primaryKey" json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint
}
