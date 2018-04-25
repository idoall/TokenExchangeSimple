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

package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/idoall/TokenExchangeSimple/info"
	"github.com/pkg/errors"
)

/**
 * //Binance 委托下单
 * @param  {[type]} as [description]
 * @return {[type]}    [description]
 //当前上下文
 ctx, _ := context.WithCancel(context.Background())
 //
 //创建一个新的服务
 as := services.NewAPIService("Binance",
	 ctx,
	 log4.NewFileLogger("access-binance-spider.log"),
	 log4.NewOutLogger())

 res, err := as.NewOrder_Binance(info.BinancePlaceRequestParams{
	 Symbol:      "BTCUSDT",
	 Quantity:    1.0,
	 Price:       29999.0,
	 Side:        internal.Binance_NewOrderRequest_SideSell,
	 TimeInForce: "GTC",
	 Type:        internal.Binance_NewOrderRequest_TypeLimit,
	 Timestamp:   time.Now(),
 })
 if err != nil {
	 fmt.Printf("main err：%v\n", err)
 } else {
	 fmt.Printf("main res:%v\n", res)
 }
*/
func (as *apiService) NewOrder_Binance(placeRequestParams info.BinancePlaceRequestParams) (*info.BinancePlaceReturn_Network, error) {

	params := make(map[string]string)
	params["symbol"] = placeRequestParams.Symbol
	params["side"] = string(placeRequestParams.Side)
	params["type"] = string(placeRequestParams.Type)
	params["timeInForce"] = string(placeRequestParams.TimeInForce)
	params["quantity"] = strconv.FormatFloat(placeRequestParams.Quantity, 'f', 5, 64)
	params["price"] = strconv.FormatFloat(placeRequestParams.Price, 'f', 5, 64)
	params["timestamp"] = strconv.FormatInt(unixMillis(placeRequestParams.Timestamp), 10)

	res, err := as.request_binance("POST", "/api/v3/order", params, true, true)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from order")
	}
	defer res.Body.Close()

	rawOrder := struct {
		Symbol        string  `json:"symbol"`
		OrderID       int64   `json:"orderId"`
		ClientOrderID string  `json:"clientOrderId"`
		TransactTime  float64 `json:"transactTime"`
	}{}
	if err := json.Unmarshal(textRes, &rawOrder); err != nil {
		return nil, errors.Wrap(err, "rawOrder unmarshal failed")
	}

	if rawOrder.OrderID == 0 {
		errMsg := struct {
			Code int64  `json:"code"`
			Msg  string `json:"msg"`
		}{}
		json.Unmarshal(textRes, &errMsg)
		return nil, errors.New(fmt.Sprintf("rawOrder-》errMsg unmarshal failed.Code:%d Msg:%s", errMsg.Code, errMsg.Msg))
	}
	fmt.Println(1111)
	t, err := timeFromUnixTimestampFloat(rawOrder.TransactTime)
	if err != nil {
		return nil, err
	}

	return &info.BinancePlaceReturn_Network{
		Symbol:        rawOrder.Symbol,
		OrderID:       rawOrder.OrderID,
		ClientOrderID: rawOrder.ClientOrderID,
		TransactTime:  t,
	}, nil
}
