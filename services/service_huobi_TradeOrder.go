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

	"github.com/idoall/TokenExchangeSimple/info"
	"github.com/pkg/errors"
)

/**
 * 火币下订单
 * @param  {[type]} as [description]
 * @return {[type]}    [description]
 // 当前上下文
 ctx, _ := context.WithCancel(context.Background())

 //创建一个新的服务
 as := services.NewAPIService("Huobi",
	 ctx,
	 log4.NewFileLogger("access-Huobi.log"),
	 log4.NewOutLogger())
 // as.GetAccounts_Huobi()
 //
 res, err := as.NewOrder_Huobi(info.HuobiPlaceRequestParams{
	 AccountID: config.Config.HuobiConfig.SPOTID,
	 Amount:    10,
	 Symbol:    "hsrusdt",
	 Price:     20,
	 Type:      internal.Huobi_NewOrderRequestType_SellLimit,
 })
 if err != nil {
	 fmt.Println(err)
 } else {
	 fmt.Println("NewOrder_Huobi Return:", res)
 }
*/
func (as *apiService) NewOrder_Huobi(placeRequestParams info.HuobiPlaceRequestParams) (*info.HuobiPlaceReturn_Network, error) {

	strRequest := "/v1/order/orders/place"
	strMethod := "POST"
	// timestamp := time.Now().UTC().Format("2006-01-02T15:04:05")
	placeReturn := info.HuobiPlaceReturn_Network{}

	mapParams := make(map[string]string)

	mapParams["account-id"] = strconv.Itoa(placeRequestParams.AccountID)
	mapParams["amount"] = strconv.FormatFloat(placeRequestParams.Amount, 'E', -1, 64)
	if placeRequestParams.Price > 0 {
		mapParams["price"] = strconv.FormatFloat(placeRequestParams.Price, 'E', -1, 64)
	}
	if 0 < len(placeRequestParams.Source) {
		mapParams["source"] = placeRequestParams.Source
	}
	mapParams["symbol"] = placeRequestParams.Symbol
	mapParams["type"] = string(placeRequestParams.Type)

	res, err := as.request_Huobi(strMethod, strRequest, mapParams, true)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from info.HuobiPlaceReturn_Network")
	}
	defer res.Body.Close()
	json.Unmarshal([]byte(textRes), &placeReturn)

	return &placeReturn, nil
}
