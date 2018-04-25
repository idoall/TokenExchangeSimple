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
	"io/ioutil"
	"strconv"

	"github.com/idoall/TokenExchangeSimple/config"
	"github.com/idoall/TokenExchangeSimple/info"
	"github.com/pkg/errors"
)

/**
 * ZB 委托下单
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

 res, err := as.NewOrder_ZB(info.ZBPlaceRequestParams{
	 Symbol: "zb_qc",
	 Amount: 10,
	 Price:  20,
	 Type:   internal.ZB_NewOrderRequestType_Sell,
 })

 if err != nil {
	 fmt.Println(err)
 } else if res.Code == 1000 {
	 fmt.Printf("操作成功:%v", res)
 } else {
	 fmt.Printf("error:%v", res)
 }
*/
func (as *apiService) NewOrder_ZB(placeRequestParams info.ZBPlaceRequestParams) (*info.ZBPlaceReturn_Network, error) {
	//设置请求参数
	mapParams := make(map[string]string)
	mapParams["method"] = "order"
	mapParams["currency"] = placeRequestParams.Symbol
	mapParams["tradeType"] = string(placeRequestParams.Type)
	if placeRequestParams.Price > 0 {
		mapParams["price"] = strconv.FormatFloat(placeRequestParams.Price, 'f', 2, 64)
	}
	if placeRequestParams.Amount > 0 {
		mapParams["amount"] = strconv.FormatFloat(placeRequestParams.Amount, 'f', 2, 64)
	}

	strRequestUrl := config.Config.ZBConfig.TRADE_URL + "/api/order"

	//发送请求
	res, err := as.request_ZB("GET", strRequestUrl, mapParams, true)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from Order.get")
	}
	defer res.Body.Close()

	//将结果解析成json
	_m := info.ZBPlaceReturn_Network{}
	if err := json.Unmarshal(textRes, &_m); err != nil {
		return nil, errors.Wrap(err, "ZBPlaceReturn_Network unmarshal failed")
	}

	return &_m, nil
}
