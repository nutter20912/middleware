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

type PagePaginator struct {
	CurrentPage *int `json:"current_page,omitempty"`
	LastPage    *int `json:"last_page,omitempty"`
	PerPage     *int `json:"per_page,omitempty"`
	Total       *int `json:"total,omitempty"`
}

type Paginator struct {
	NextCursor *string `json:"next_cursor,omitempty"`
}

type PositionStream struct {
	Open   []*SpotPosition       `json:"open"`
	Closed []*SpotPositionClosed `json:"closed"`
}

type Posts struct {
	Data      []*Post    `json:"data"`
	Paginator *Paginator `json:"paginator,omitempty"`
}

type TradeStream struct {
	AggTrade *AggTradeData `json:"agg_trade"`
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
