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
	"github.com/idoall/TokenExchangeSimple/internal"
	"github.com/pkg/errors"
)

/**
 * 获取 火币的 Kline
 * @param  {[type]} as [description]
 * @return {[type]}    [description]
 *
 //当前上下文
 ctx, _ := context.WithCancel(context.Background())

 //创建一个新的服务
 as := services.NewAPIService("Huobi",
	 ctx,
	 log4.NewFileLogger("access-binance-spider.log"),
	 log4.NewOutLogger())

 //查询  火币 K线图
 res, err := as.Klines_Huobi(info.HuobiKlinesRequestParams{
	 Symbol: "btcusdt",
	 Period: internal.Huobi_Interval_Hour,
	 Size:   100,
 })
 if err != nil {
	 fmt.Println(err)
 }
 fmt.Println(res)
*/
func (as *apiService) Klines_Huobi(kr info.HuobiKlinesRequestParams) (*info.HuobiKLineReturn_Network, error) {
	kLineReturn := info.HuobiKLineReturn_Network{}

	mapParams := make(map[string]string)
	mapParams["symbol"] = kr.Symbol
	mapParams["period"] = string(internal.Huobi_Interval_Hour)
	mapParams["size"] = strconv.Itoa(kr.Size)

	res, err := as.request_Huobi("GET", "/market/history/kline", mapParams, false)
	if err != nil {
		return &info.HuobiKLineReturn_Network{}, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &info.HuobiKLineReturn_Network{}, errors.Wrap(err, "unable to read response from info.HuobiKLineReturn_Network")
	}
	defer res.Body.Close()

	json.Unmarshal([]byte(textRes), &kLineReturn)
	return &kLineReturn, nil
}
