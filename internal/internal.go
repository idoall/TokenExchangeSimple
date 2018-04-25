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

package internal

/**
 *  火币交易类型
 * @type {[type]}
 */
type Huobi_NewOrderRequestParamsType string

var (
	/**
	 * 市价买
	 * @type {[type]}
	 */
	Huobi_NewOrderRequestType_BuyMarkdt = Huobi_NewOrderRequestParamsType("buy-market")
	/**
	 * 市价卖
	 * @type {[type]}
	 */
	Huobi_NewOrderRequestType_SellMarkdt = Huobi_NewOrderRequestParamsType("sell-market")
	/**
	 * 限价买
	 * @type {[type]}
	 */
	Huobi_NewOrderRequestType_BuyLimit = Huobi_NewOrderRequestParamsType("buy-limit")
	/**
	 * 限价卖
	 * @type {[type]}
	 */
	Huobi_NewOrderRequestType_SellLimit = Huobi_NewOrderRequestParamsType("sell-limit")
)

/**
 * ZB 交易类型
 * @type {[type]}
 */
type ZB_NewOrderRequestParamsType string

var (
	/**
	 * 买
	 * @type {[type]}
	 */
	ZB_NewOrderRequestType_Buy = ZB_NewOrderRequestParamsType("1")
	/**
	 * 卖
	 * @type {[type]}
	 */
	ZB_NewOrderRequestType_Sell = ZB_NewOrderRequestParamsType("0")
)

/**
 * Binance 交易类型
 * @type {[type]}
 */
type Binance_NewOrderRequestParamsSide string

var (
	/**
	 * 买
	 * @type {[type]}
	 */
	Binance_NewOrderRequest_SideBuy = Binance_NewOrderRequestParamsSide("BUY")
	/**
	 * 卖
	 * @type {[type]}
	 */
	Binance_NewOrderRequest_SideSell = Binance_NewOrderRequestParamsSide("SELL")
)

/**
 * Binance 交易类型
 * @type {[type]}
 */
type Binance_NewOrderRequestParamsOrder string

var (
	Binance_NewOrderRequest_TypeLimit  = Binance_NewOrderRequestParamsOrder("LIMIT")  // 限价
	Binance_NewOrderRequest_TypeMarket = Binance_NewOrderRequestParamsOrder("MARKET") // 市场价
)

// Interval represents interval enum.
type Huobi_Interval string

var (
	Huobi_Interval_Minute         = Huobi_Interval("1min")
	Huobi_Interval_FiveMinutes    = Huobi_Interval("5min")
	Huobi_Interval_FifteenMinutes = Huobi_Interval("15min")
	Huobi_Interval_ThirtyMinutes  = Huobi_Interval("30min")
	Huobi_Interval_Hour           = Huobi_Interval("60min")
	Huobi_Interval_Day            = Huobi_Interval("1day")
	Huobi_Interval_Week           = Huobi_Interval("1week")
	Huobi_Interval_Mohth          = Huobi_Interval("1mon")
	Huobi_Interval_Year           = Huobi_Interval("1year")
)

// Interval represents interval enum.
type Binance_Interval string

var (
	Binance_Interval_Minute         = Binance_Interval("1m")
	Binance_Interval_ThreeMinutes   = Binance_Interval("3m")
	Binance_Interval_FiveMinutes    = Binance_Interval("5m")
	Binance_Interval_FifteenMinutes = Binance_Interval("15m")
	Binance_Interval_ThirtyMinutes  = Binance_Interval("30m")
	Binance_Interval_Hour           = Binance_Interval("1h")
	Binance_Interval_TwoHours       = Binance_Interval("2h")
	Binance_Interval_FourHours      = Binance_Interval("4h")
	Binance_Interval_SixHours       = Binance_Interval("6h")
	Binance_Interval_EightHours     = Binance_Interval("8h")
	Binance_Interval_TwelveHours    = Binance_Interval("12h")
	Binance_Interval_Day            = Binance_Interval("1d")
	Binance_Interval_ThreeDays      = Binance_Interval("3d")
	Binance_Interval_Week           = Binance_Interval("1w")
	Binance_Interval_Month          = Binance_Interval("1M")
)

// Interval represents interval enum.
type ZB_Interval string

var (
	ZB_Interval_Minute         = ZB_Interval("1min")
	ZB_Interval_ThreeMinutes   = ZB_Interval("3min")
	ZB_Interval_FiveMinutes    = ZB_Interval("5min")
	ZB_Interval_FifteenMinutes = ZB_Interval("15min")
	ZB_Interval_ThirtyMinutes  = ZB_Interval("30min")
	ZB_Interval_Hour           = ZB_Interval("1hour")
	ZB_Interval_TwoHours       = ZB_Interval("2hour")
	ZB_Interval_FourHours      = ZB_Interval("4hour")
	ZB_Interval_SixHours       = ZB_Interval("6hour")
	ZB_Interval_TwelveHours    = ZB_Interval("12hour")
	ZB_Interval_Day            = ZB_Interval("1day")
	ZB_Interval_ThreeDays      = ZB_Interval("3day")
	ZB_Interval_Week           = ZB_Interval("1week")
)
