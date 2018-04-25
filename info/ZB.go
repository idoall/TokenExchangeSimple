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

//-------------Accounts

type ZBAccountsReturn_Coins_Network struct {
	Freez       float64 `json:"freez"`       //冻结资产
	EnName      string  `json:"enName"`      //币种英文名
	UnitDecimal int     `json:"unitDecimal"` //保留小数位
	UnName      string  `json:"cnName"`      //币种中文名
	UnitTag     string  `json:"unitTag"`     //币种符号
	Available   float64 `json:"available"`   //可用资产
	Key         string  `json:"key"`         //币种
}

type ZBAccountsReturn_Network struct {
	UserName               string                            `json:"username"`               //用户名
	Trade_Password_Enabled bool                              `json:"trade_password_enabled"` //是否开通交易密码
	Auth_Google_Enabled    bool                              `json:"auth_google_enabled"`    //是否开通谷歌验证
	Auth_Mobile_Enabled    bool                              `json:"auth_mobile_enabled"`    //是否开通手机验证
	Coins                  []*ZBAccountsReturn_Coins_Network `json:"list"`                   // 资产列表
}

//-------------Ticket

type ZBTicket_Network struct {
	Date   string                 `json:"date"`
	Ticket zBTicket_Child_Network `json:"ticker"`
}

type zBTicket_Child_Network struct {
	Vol  string //成交量(最近的24小时)
	Last string //最新成交价
	Sell string //卖一价
	Buy  string //买一价
	High string //最高价
	Low  string //最低价
}

//-------------Trade Place Order

type ZBPlaceRequestParams struct {
	Amount float64                               `json:"amount"`    // 交易数量
	Price  float64                               `json:"price"`     // 下单价格,
	Symbol string                                `json:"currency"`  // 交易对, btcusdt, bccbtc......
	Type   internal.ZB_NewOrderRequestParamsType `json:"tradeType"` // 订单类型, buy-market: 市价买, sell-market: 市价卖, buy-limit: 限价买, sell-limit: 限价卖
}

type ZBPlaceReturn_Network struct {
	Code    int    `json:"code"`    //返回代码
	Message string `json:"message"` //提示信息
	ID      string `json:"id"`      //委托挂单号
}

//-------------Kline

// KlinesRequest represents Klines request data.
type ZBKlinesRequestParams struct {
	Market string               //交易对, zb_qc,zb_usdt,zb_btc...
	Type   internal.ZB_Interval //K线类型, 1min, 3min, 15min, 30min, 1hour......
	Since  string               //从这个时间戳之后的
	Size   int                  //返回数据的条数限制(默认为1000，如果返回数据多于1000条，那么只返回1000条)
}

type ZBKLineData_Network struct {
	ID        float64   `json:"id"` // K线ID
	KlineTime time.Time `json:"klineTime"`
	Open      float64   `json:"open"`  // 开盘价
	Close     float64   `json:"close"` // 收盘价, 当K线为最晚的一根时, 时最新成交价
	Low       float64   `json:"low"`   // 最低价
	High      float64   `json:"high"`  // 最高价
	Volume    float64   `json:"vol"`   // 成交量
}

type ZBKLineReturn_Network struct {
	// Data      string                `json:"data"`      // 买入货币
	MoneyType string                 `json:"moneyType"` // 卖出货币
	Symbol    string                 `json:"symbol"`    // 内容说明
	Data      []*ZBKLineData_Network `json:"data"`      // KLine数据
}
