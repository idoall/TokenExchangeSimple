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
 * [func description]
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

 //查询 ZB Kline 数据
 d, _ := time.ParseDuration("-1000h")
 _since := time.Now().Add(d).Unix() * int64(time.Microsecond)
 res, err := as.Klines_ZB(info.ZBKlinesRequestParams{
	 Market: "zb_qc",
	 Type:   internal.ZB_Interval_Hour,
	 Since:  strconv.FormatInt(_since, 10),
	 Size:   100,
 })
 if err != nil {
	 fmt.Println(err)
 } else {
	 fmt.Printf("%v", res)
 }
*/
func (as *apiService) Klines_ZB(kr info.ZBKlinesRequestParams) (*info.ZBKLineReturn_Network, error) {

	strRequest := "/data/v1/kline"

	mapParams := make(map[string]string)
	mapParams["market"] = kr.Market
	mapParams["type"] = string(kr.Type)
	if kr.Since != "" {
		mapParams["since"] = kr.Since
	}
	mapParams["size"] = strconv.Itoa(kr.Size)

	strRequestUrl := config.Config.ZBConfig.MARKET_URL + strRequest

	res, err := as.request_ZB("GET", strRequestUrl, mapParams, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from kline.get")
	}
	defer res.Body.Close()

	var rawKlines map[string]interface{}
	if err := json.Unmarshal([]byte(textRes), &rawKlines); err != nil {
		return nil, errors.Wrap(err, "rawKlines unmarshal failed")
	}
	if rawKlines == nil || rawKlines["symbol"] == nil {
		return nil, errors.Wrap(err, "rawKlines is nil")
	}
	// fmt.Println(rawKlines)

	_m := &info.ZBKLineReturn_Network{
		Symbol:    rawKlines["symbol"].(string),
		MoneyType: rawKlines["moneyType"].(string),
	}

	//对于 Data数据，再次解析
	_rawKlineDatasString, _ := json.Marshal(rawKlines["data"].([]interface{}))
	rawKlineDatas := [][]interface{}{}
	if err := json.Unmarshal(_rawKlineDatasString, &rawKlineDatas); err != nil {
		return nil, errors.Wrap(err, "rawKlines unmarshal failed")
	}

	for _, k := range rawKlineDatas {
		// s := strconv.FormatFloat(k[0].(float64), 'E', -1, 64)
		//time.Unix(_item.Timestamp, 0).Format("2006-01-02 15:04:05")
		ot, err := timeFromUnixTimestampFloat(k[0])
		if err != nil {
			return nil, errors.Wrap(err, "cannot parse Kline.OpenTime")
		}
		_m.Data = append(_m.Data, &info.ZBKLineData_Network{
			ID:        k[0].(float64),
			KlineTime: ot,
			Open:      k[1].(float64),
			High:      k[2].(float64),
			Low:       k[3].(float64),
			Close:     k[4].(float64),
			Volume:    k[5].(float64),
		})
	}
	return _m, nil
}
