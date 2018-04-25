// Copyright 2016 mshk.top, lion@mshk.top
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package info

import (
	"time"

	"github.com/idoall/TokenExchangeSimple/internal"
)

//------------Kline

// BinanceKlinesRequestParams represents Klines request data.
type BinanceKlinesRequestParams struct {
	Symbol    string                    //必填项，交易对:LTCBTC,BTCUSDT
	Interval  internal.Binance_Interval //查询时间段
	Limit     int                       // Default 500; max 500.
	StartTime int64
	EndTime   int64
}

//binance网络上的Kline解析
type BinanceKline_Network struct {
	OpenTime                 time.Time
	Open                     float64
	High                     float64
	Low                      float64
	Close                    float64
	Volume                   float64
	CloseTime                time.Time
	QuoteAssetVolume         float64
	NumberOfTrades           int
	TakerBuyBaseAssetVolume  float64
	TakerBuyQuoteAssetVolume float64
}

//-----------Account

// AccountRequest represents Account request data.
type BinanceAccountRequest_Network struct {
	RecvWindow time.Duration
	Timestamp  time.Time
}

// Account represents user's account information.
type BinanceAccount_Network struct {
	MakerCommision  int64
	TakerCommision  int64
	BuyerCommision  int64
	SellerCommision int64
	CanTrade        bool
	CanWithdraw     bool
	CanDeposit      bool
	Balances        []*BinanceBalance_Network
}

// Balance groups balance-related information.
type BinanceBalance_Network struct {
	Asset  string
	Free   float64
	Locked float64
}

//----------Trade Order
// BinancePlaceRequestParams represents NewOrder request data.
type BinancePlaceRequestParams struct {
	Symbol           string
	Side             internal.Binance_NewOrderRequestParamsSide
	Type             internal.Binance_NewOrderRequestParamsOrder
	TimeInForce      string
	Quantity         float64 //数量
	Price            float64 //单价
	NewClientOrderID string
	StopPrice        float64
	IcebergQty       float64
	Timestamp        time.Time
}

// BinancePlaceReturn_Network represents data from processed order.
type BinancePlaceReturn_Network struct {
	Symbol        string
	OrderID       int64
	ClientOrderID string
	TransactTime  time.Time
}
