package model

import (
	"encoding/json"
	orderV1 "middleware/proto/order/v1"
	"strings"
)

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  string `json:"user_id"`
	User    *User  `json:"user"`
}

type Comment struct {
	ID      string `json:"id"`
	UserId  string `json:"user_id"`
	Content string `json:"content"`
	User    *User  `json:"user"`
}

type DepositOrder struct {
	ID        string               `json:"id"`
	UserID    string               `json:"user_id"`
	Status    DepositStatus        `json:"status"`
	Amount    float64              `json:"amount"`
	Memo      string               `json:"memo"`
	CreatedAt string               `json:"created_at"`
	UpdatedAt string               `json:"updated_at"`
	Events    []*DepositOrderEvent `json:"events"`
}

type SpotPositionClosed struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	CreatedAt    string    `json:"created_at"`
	Symbol       string    `json:"symbol"`
	Side         OrderSide `json:"side"`
	Quantity     float64   `json:"quantity"`
	OpenOrderID  string    `json:"open_order_id"`
	OpenPrice    float64   `json:"open_price"`
	OpenFee      float64   `json:"open_fee"`
	CloseOrderID string    `json:"close_order_id"`
	ClosePrice   float64   `json:"close_price"`
	CloseFee     float64   `json:"close_fee"`
}

func (s *SpotPositionClosed) UnmarshalJSON(data []byte) error {
	var temp orderV1.SpotPositionClosed

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	s.ID = temp.Id
	s.UserID = temp.UserId
	s.CreatedAt = temp.CreatedAt
	s.Symbol = temp.Symbol
	s.Quantity = temp.Quantity
	s.OpenOrderID = temp.OpenOrderId
	s.OpenPrice = temp.OpenPrice
	s.OpenFee = temp.OpenFee
	s.CloseOrderID = temp.CloseOrderId
	s.ClosePrice = temp.ClosePrice
	s.CloseFee = temp.CloseFee

	s.Side = OrderSide(
		strings.TrimPrefix(temp.Side.String(), "ORDER_SIDE_"),
	)

	return nil
}

type SpotPosition struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
	Symbol       string    `json:"symbol"`
	Side         OrderSide `json:"side"`
	Quantity     float64   `json:"quantity"`
	OrderID      string    `json:"order_id"`
	Price        float64   `json:"price"`
	Fee          float64   `json:"fee"`
	OpenQuantity float64   `json:"open_quantity"`
}

func (s *SpotPosition) UnmarshalJSON(data []byte) error {
	var temp orderV1.SpotPosition

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	s.ID = temp.Id
	s.UserID = temp.UserId
	s.CreatedAt = temp.CreatedAt
	s.UpdatedAt = temp.UpdatedAt
	s.Symbol = temp.Symbol
	s.Quantity = temp.Quantity
	s.OrderID = temp.OrderId
	s.Price = temp.Price
	s.Fee = temp.Fee
	s.OpenQuantity = temp.OpenQuantity

	s.Side = OrderSide(
		strings.TrimPrefix(temp.Side.String(), "ORDER_SIDE_"),
	)

	return nil
}

type AggTradeData struct {
	EventType       string  `json:"event_type"`
	EventTime       uint64  `json:"event_time"`
	Symbol          string  `json:"symbol"`
	Price           float64 `json:"price"`
	Quantity        float64 `json:"quantity"`
	TransactionTime uint64  `json:"transaction_time"`
	IsSell          bool    `json:"is_sell"`
}
