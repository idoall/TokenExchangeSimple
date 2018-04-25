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
	"github.com/idoall/TokenExchangeSimple/internal"
)

//-------------Accounts

type HuobiAccountsData_Network struct {
	ID     int64  `json:"id"`      // Account ID
	Type   string `json:"type"`    // 账户类型, spot: 现货账户
	State  string `json:"state"`   // 账户状态, working: 正常, lock: 账户被锁定
	UserID int64  `json:"user-id"` // 用户ID
}

type HuobiAccountsReturn_Network struct {
	Status  string                      `json:"status"` // 请求状态
	Data    []HuobiAccountsData_Network `json:"data"`   // 用户数据
	ErrCode string                      `json:"err-code"`
	ErrMsg  string                      `json:"err-msg"`
}

//-------------Trade Place Order

type HuobiCheckTradeOrderParams struct {
	Symbol string //交易对
	Period int    //默认计算几天的KDJ一般都是9天的数据
	RSV    []float64
	K      []float64
	D      []float64
	J      []float64
}

type HuobiPlaceRequestParams struct {
	AccountID int                                      `json:"account-id"` // 账户 ID，使用accounts方法获得。币币交易使用‘spot’账户的accountid；借贷资产交易，请使用‘margin’账户的accountid
	Amount    float64                                  `json:"amount"`     // 限价表示下单数量, 市价买单时表示买多少钱, 市价卖单时表示卖多少币
	Price     float64                                  `json:"price"`      // 下单价格, 市价单不传该参数
	Source    string                                   `json:"source"`     // 订单来源, api: API调用, margin-api: 借贷资产交易
	Symbol    string                                   `json:"symbol"`     // 交易对, btcusdt, bccbtc......
	Type      internal.Huobi_NewOrderRequestParamsType `json:"type"`       // 订单类型, buy-market: 市价买, sell-market: 市价卖, buy-limit: 限价买, sell-limit: 限价卖
}

type HuobiPlaceReturn_Network struct {
	Status  string `json:"status"`
	Data    string `json:"data"`
	ErrCode string `json:"err-code"`
	ErrMsg  string `json:"err-msg"`
}

//-------------Trade

type HuobiTradeData_Network struct {
	ID        int64   `json:"id"`        //成交ID
	Price     float64 `json:"price"`     // 成交价
	Amount    float64 `json:"amount"`    // 成交量
	Direction string  `json:"direction"` // 主动成交方向
	Ts        int64   `json:"ts"`        // 成交时间
}

type HuobiTradeTick_Network struct {
	ID   int64                    `json:"id"`   // 消息ID
	Ts   int64                    `json:"ts"`   // 最新成交时间
	Data []HuobiTradeData_Network `json:"data"` // Trade数据
}

type HuobiTradeReturn_Network struct {
	Status  string                   `json:"status"` // 请求状态, ok或者error
	Ch      string                   `json:"ch"`     // 数据所属的Channel, 格式: market.$symbol.trade.detail
	Ts      int64                    `json:"ts"`     // 发送时间
	Data    []HuobiTradeTick_Network `json:"data"`   // 成交记录
	ErrCode string                   `json:"err-code"`
	ErrMsg  string                   `json:"err-msg"`
}

//-------------Kline

// KlinesRequest represents Klines request data.
type HuobiKlinesRequestParams struct {
	Symbol string                  //交易对, btcusdt, bccbtc......
	Period internal.Huobi_Interval //K线类型, 1min, 5min, 15min......
	Size   int                     //获取数量, [1-2000]
}

type HuobiKLineData_Network struct {
	ID     int64   `json:"id"`     // K线ID
	Amount float64 `json:"amount"` // 成交量
	Count  int64   `json:"count"`  // 成交笔数
	Open   float64 `json:"open"`   // 开盘价
	Close  float64 `json:"close"`  // 收盘价, 当K线为最晚的一根时, 时最新成交价
	Low    float64 `json:"low"`    // 最低价
	High   float64 `json:"high"`   // 最高价
	Vol    float64 `json:"vol"`    // 成交额, 即SUM(每一笔成交价 * 该笔的成交数量)
}

type HuobiKLineReturn_Network struct {
	Status  string                   `json:"status"`   // 请求处理结果, "ok"、"error"
	Ts      int64                    `json:"ts"`       // 响应生成时间点, 单位毫秒
	Data    []HuobiKLineData_Network `json:"data"`     // KLine数据
	Ch      string                   `json:"ch"`       // 数据所属的Channel, 格式: market.$symbol.kline.$period
	ErrCode string                   `json:"err-code"` // 错误代码
	ErrMsg  string                   `json:"err-msg"`  // 错误提示
}
