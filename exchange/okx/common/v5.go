package common

import (
	"encoding/json"
	"okx-bot/exchange/options"
)

type OKxV5 struct {
	UriOpts       options.UriOptions
	UnmarshalOpts options.UnmarshalerOptions
}

type BaseResp struct {
	Code int             `json:"code,string"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

func New() *OKxV5 {
	unmarshaler := new(RespUnmarshaler)

	f := &OKxV5{
		UriOpts: options.UriOptions{
			Endpoint:            "https://www.okx.com",
			KlineUri:            "/api/v5/market/candles",
			TickerUri:           "/api/v5/market/ticker",
			DepthUri:            "/api/v5/market/books",
			NewOrderUri:         "/api/v5/trade/order",
			GetOrderUri:         "/api/v5/trade/order",
			GetHistoryOrdersUri: "/api/v5/trade/orders-history",
			GetPendingOrdersUri: "/api/v5/trade/orders-pending",
			CancelOrderUri:      "/api/v5/trade/cancel-order",
			GetAccountUri:       "/api/v5/account/balance",
			GetPositionsUri:     "/api/v5/account/positions",
			GetExchangeInfoUri:  "/api/v5/public/instruments",
		},
		UnmarshalOpts: options.UnmarshalerOptions{
			ResponseUnmarshaler:                  unmarshaler.UnmarshalResponse,
			KlineUnmarshaler:                     unmarshaler.UnmarshalGetKlineResponse,
			TickerUnmarshaler:                    unmarshaler.UnmarshalTicker,
			DepthUnmarshaler:                     unmarshaler.UnmarshalDepth,
			CreateOrderResponseUnmarshaler:       unmarshaler.UnmarshalCreateOrderResponse,
			GetPendingOrdersResponseUnmarshaler:  unmarshaler.UnmarshalGetPendingOrdersResponse,
			GetHistoryOrdersResponseUnmarshaler:  unmarshaler.UnmarshalGetHistoryOrdersResponse,
			CancelOrderResponseUnmarshaler:       unmarshaler.UnmarshalCancelOrderResponse,
			GetOrderInfoResponseUnmarshaler:      unmarshaler.UnmarshalGetOrderInfoResponse,
			GetAccountResponseUnmarshaler:        unmarshaler.UnmarshalGetAccountResponse,
			GetPositionsResponseUnmarshaler:      unmarshaler.UnmarshalGetPositionsResponse,
			GetFuturesAccountResponseUnmarshaler: unmarshaler.UnmarshalGetFuturesAccountResponse,
			GetExchangeInfoResponseUnmarshaler:   unmarshaler.UnmarshalGetExchangeInfoResponse,
		},
	}

	return f
}

func (okx *OKxV5) WithUriOption(opts ...options.UriOption) *OKxV5 {
	for _, opt := range opts {
		opt(&okx.UriOpts)
	}
	return okx
}

func (okx *OKxV5) WithUnmarshalOption(opts ...options.UnmarshalerOption) *OKxV5 {
	for _, opt := range opts {
		opt(&okx.UnmarshalOpts)
	}
	return okx
}

func (okx *OKxV5) NewPrvApi(opts ...options.ApiOption) *Prv {
	api := NewPrvApi(opts...)
	api.OKxV5 = okx
	return api
}
