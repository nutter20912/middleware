// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Comments struct {
	Data      []*Comment `json:"data"`
	Paginator *Paginator `json:"paginator,omitempty"`
}

type DepositOrderEvent struct {
	UserID  string        `json:"user_id"`
	OrderID string        `json:"order_id"`
	Status  DepositStatus `json:"status"`
	Amount  float64       `json:"amount"`
	Memo    string        `json:"memo"`
	Time    string        `json:"time"`
}

type DepthData struct {
	Bids [][]string `json:"bids"`
	Asks [][]string `json:"asks"`
}

type Kline struct {
	StartTime uint64 `json:"start_time"`
	EndTime   uint64 `json:"end_time"`
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	Open      string `json:"open"`
	Close     string `json:"close"`
	High      string `json:"high"`
	Low       string `json:"low"`
}

type KlineData struct {
	EventType string `json:"event_type"`
	EventTime uint64 `json:"event_time"`
	Symbol    string `json:"symbol"`
	Kline     *Kline `json:"kline"`
}

type PagePaginator struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
}

type Paginator struct {
	NextCursor *string `json:"next_cursor,omitempty"`
}

type PositionStream struct {
	Open []*SpotPosition `json:"open"`
}

type Posts struct {
	Data      []*Post    `json:"data"`
	Paginator *Paginator `json:"paginator,omitempty"`
}

type SpotOrder struct {
	ID        string      `json:"id"`
	UserID    string      `json:"user_id"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	Symbol    string      `json:"symbol"`
	Price     float64     `json:"price"`
	Quantity  float64     `json:"quantity"`
	Side      OrderSide   `json:"side"`
	Type      OrderType   `json:"type"`
	Status    OrderStatus `json:"status"`
	Memo      string      `json:"memo"`
}

type SpotOrderEvent struct {
	UserID   string      `json:"user_id"`
	OrderID  string      `json:"order_id"`
	Time     string      `json:"time"`
	Symbol   string      `json:"symbol"`
	Quantity float64     `json:"quantity"`
	Side     OrderSide   `json:"side"`
	Type     OrderType   `json:"type"`
	Status   OrderStatus `json:"status"`
	Price    float64     `json:"price"`
	Memo     string      `json:"memo"`
}

type SpotOrderFilter struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type SpotOrders struct {
	Data      []*SpotOrder   `json:"data"`
	Paginator *PagePaginator `json:"paginator,omitempty"`
}

type SpotPositionClosedFilter struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Symbol    string `json:"symbol"`
}

type SpotPositionCloseds struct {
	Data      []*SpotPositionClosed `json:"data"`
	Paginator *PagePaginator        `json:"paginator,omitempty"`
}

type SpotPositions struct {
	Data      []*SpotPosition `json:"data"`
	Paginator *PagePaginator  `json:"paginator,omitempty"`
}

type TradeStream struct {
	AggTrade *AggTradeData `json:"agg_trade,omitempty"`
	Kline    *KlineData    `json:"kline,omitempty"`
	Depth    *DepthData    `json:"depth,omitempty"`
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type Wallet struct {
	Amount    float64 `json:"amount"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type WalletEvent struct {
	ID      string          `json:"id"`
	UserID  string          `json:"user_id"`
	OrderID string          `json:"order_id"`
	Type    WalletEventType `json:"type"`
	Time    string          `json:"time"`
	Change  float64         `json:"change"`
	Memo    string          `json:"memo"`
}

type WalletEventFilter struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type WalletEvents struct {
	Data      []*WalletEvent `json:"data"`
	Paginator *PagePaginator `json:"paginator,omitempty"`
}

type WalletStream struct {
	Info   *Wallet        `json:"info,omitempty"`
	Events []*WalletEvent `json:"events,omitempty"`
}

type DepositStatus string

const (
	DepositStatusUnspecified DepositStatus = "UNSPECIFIED"
	DepositStatusPending     DepositStatus = "PENDING"
	DepositStatusProcessing  DepositStatus = "PROCESSING"
	DepositStatusCompleted   DepositStatus = "COMPLETED"
	DepositStatusFailed      DepositStatus = "FAILED"
	DepositStatusCanceled    DepositStatus = "CANCELED"
)

var AllDepositStatus = []DepositStatus{
	DepositStatusUnspecified,
	DepositStatusPending,
	DepositStatusProcessing,
	DepositStatusCompleted,
	DepositStatusFailed,
	DepositStatusCanceled,
}

func (e DepositStatus) IsValid() bool {
	switch e {
	case DepositStatusUnspecified, DepositStatusPending, DepositStatusProcessing, DepositStatusCompleted, DepositStatusFailed, DepositStatusCanceled:
		return true
	}
	return false
}

func (e DepositStatus) String() string {
	return string(e)
}

func (e *DepositStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DepositStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DepositStatus", str)
	}
	return nil
}

func (e DepositStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderSide string

const (
	OrderSideUnspecified OrderSide = "UNSPECIFIED"
	OrderSideBuy         OrderSide = "BUY"
	OrderSideSell        OrderSide = "SELL"
)

var AllOrderSide = []OrderSide{
	OrderSideUnspecified,
	OrderSideBuy,
	OrderSideSell,
}

func (e OrderSide) IsValid() bool {
	switch e {
	case OrderSideUnspecified, OrderSideBuy, OrderSideSell:
		return true
	}
	return false
}

func (e OrderSide) String() string {
	return string(e)
}

func (e *OrderSide) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderSide(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderSide", str)
	}
	return nil
}

func (e OrderSide) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderStatus string

const (
	OrderStatusUnspecified OrderStatus = "UNSPECIFIED"
	OrderStatusNew         OrderStatus = "NEW"
	OrderStatusPenn        OrderStatus = "PENN"
	OrderStatusCanceled    OrderStatus = "CANCELED"
	OrderStatusFilled      OrderStatus = "FILLED"
	OrderStatusRejected    OrderStatus = "REJECTED"
)

var AllOrderStatus = []OrderStatus{
	OrderStatusUnspecified,
	OrderStatusNew,
	OrderStatusPenn,
	OrderStatusCanceled,
	OrderStatusFilled,
	OrderStatusRejected,
}

func (e OrderStatus) IsValid() bool {
	switch e {
	case OrderStatusUnspecified, OrderStatusNew, OrderStatusPenn, OrderStatusCanceled, OrderStatusFilled, OrderStatusRejected:
		return true
	}
	return false
}

func (e OrderStatus) String() string {
	return string(e)
}

func (e *OrderStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderStatus", str)
	}
	return nil
}

func (e OrderStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderType string

const (
	OrderTypeUnspecified OrderType = "UNSPECIFIED"
	OrderTypeLimit       OrderType = "LIMIT"
	OrderTypeMarket      OrderType = "MARKET"
)

var AllOrderType = []OrderType{
	OrderTypeUnspecified,
	OrderTypeLimit,
	OrderTypeMarket,
}

func (e OrderType) IsValid() bool {
	switch e {
	case OrderTypeUnspecified, OrderTypeLimit, OrderTypeMarket:
		return true
	}
	return false
}

func (e OrderType) String() string {
	return string(e)
}

func (e *OrderType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderType", str)
	}
	return nil
}

func (e OrderType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type WalletEventType string

const (
	WalletEventTypeUnspecified WalletEventType = "UNSPECIFIED"
	WalletEventTypeSystem      WalletEventType = "SYSTEM"
	WalletEventTypeDeposit     WalletEventType = "DEPOSIT"
	WalletEventTypeWithdraw    WalletEventType = "WITHDRAW"
	WalletEventTypeSpotOrder   WalletEventType = "SPOT_ORDER"
)

var AllWalletEventType = []WalletEventType{
	WalletEventTypeUnspecified,
	WalletEventTypeSystem,
	WalletEventTypeDeposit,
	WalletEventTypeWithdraw,
	WalletEventTypeSpotOrder,
}

func (e WalletEventType) IsValid() bool {
	switch e {
	case WalletEventTypeUnspecified, WalletEventTypeSystem, WalletEventTypeDeposit, WalletEventTypeWithdraw, WalletEventTypeSpotOrder:
		return true
	}
	return false
}

func (e WalletEventType) String() string {
	return string(e)
}

func (e *WalletEventType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WalletEventType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WalletEventType", str)
	}
	return nil
}

func (e WalletEventType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
