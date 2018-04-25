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

	"github.com/pkg/errors"
)

/**
 * 获取最新价格
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

 res, err := as.GetTradeNowPrice_Binance("BTCUSDT")
 if err != nil {
	 fmt.Println(err)
 } else {
	 fmt.Printf("%v", res)
 }
*/
func (as *apiService) GetTradeNowPrice_Binance(strSymbol string) (float64, error) {
	params := make(map[string]string)
	params["symbol"] = strSymbol

	res, err := as.request_binance("GET", "/api/v3/ticker/price", params, false, false)

	if err != nil {
		return 0, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, errors.Wrap(err, "unable to read response from price.get")
	}
	defer res.Body.Close()

	rawRow := struct {
		Symbol string `json:"symbol"`
		Price  string `json:"price"`
	}{}
	if err := json.Unmarshal(textRes, &rawRow); err != nil {
		return 0, errors.Wrap(err, "rawRow unmarshal failed")
	}

	price, err := strconv.ParseFloat(rawRow.Price, 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}
