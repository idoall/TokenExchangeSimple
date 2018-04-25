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

package main

import (
	"context"
	"fmt"

	"github.com/idoall/TokenExchangeSimple/config"
	"github.com/idoall/TokenExchangeSimple/log4"
	"github.com/idoall/TokenExchangeSimple/services"
)

func main() {

	//创建一个新的服务
	as := createService("mshk.top", "access.log")

	//查询 ZB 帐号信息
	res, err := as.GetAccounts_ZB()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v", res)
	}

	// res, err := as.NewOrder_ZB(info.ZBPlaceRequestParams{
	// 	Symbol: "btc_qc",
	// 	Amount: 0.5,
	// 	Price:  90000,
	// 	Type:   internal.ZB_NewOrderRequestType_Sell,
	// })
	//
	// if err != nil {
	// 	fmt.Println(err)
	// } else if res.Code == 1000 {
	// 	fmt.Printf("操作成功:%v", res)
	// } else {
	// 	fmt.Printf("error:%v", res)
	// }

	//获取 火币 帐号ID
	// as.GetAccounts_Huobi()

	//火币下订单
	// res, err := as.NewOrder_Huobi(info.HuobiPlaceRequestParams{
	// 	AccountID: config.Config.HuobiConfig.SPOTID,
	// 	Amount:    0.5,
	// 	Symbol:    "btcusdt",
	// 	Price:     20000,
	// 	Type:      internal.Huobi_NewOrderRequestType_SellLimit,
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("NewOrder_Huobi Return:", res)
	// }

	//币安获取 帐户信息
	// res, err := as.GetAccounts_Binance(info.BinanceAccountRequest_Network{
	// 	RecvWindow: 5 * time.Second,
	// 	Timestamp:  time.Now(),
	// })
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// } else {
	// 	fmt.Printf("%v\n", res)
	// }

	//币安下订单
	// newOrder_Binance, err := as.NewOrder_Binance(info.BinancePlaceRequestParams{
	// 	Symbol:      "BTCUSDT",
	// 	Quantity:    0.5,
	// 	Price:       29999.0,
	// 	Side:        internal.Binance_NewOrderRequest_SideSell,
	// 	TimeInForce: "GTC",
	// 	Type:        internal.Binance_NewOrderRequest_TypeLimit,
	// 	Timestamp:   time.Now(),
	// })
	// if err != nil {
	// 	fmt.Printf("main err：%v\n", err)
	// } else {
	// 	fmt.Printf("main res:%v\n", newOrder_Binance)
	// }

}

//create service
func createService(appName, logFileName string) services.Service {
	//当前上下文
	ctx, _ := context.WithCancel(context.Background())

	//创建一个新的服务
	as := services.NewAPIService(appName,
		ctx,
		log4.NewFileLogger(logFileName),
		log4.NewOutLogger(),
	)

	return as
}

/**
 * 初始化，在main之前调用
 */
func init() {

	fmt.Println("Init Config ......")
	config.Initialization()
	fmt.Println("Init Config Done ......")
}
