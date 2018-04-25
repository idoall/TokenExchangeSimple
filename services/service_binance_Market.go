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

//------------------------------------------------------------------------------------------
// 行情API

/**
* 获取K线数据
* @symbol: 交易对, BNBBTC,LTCBTC,ADAUSDT......
* @interval: K线类型, 1m,3m,5m,15m,30m,1h,2h,4h,6h,8h,12h,1d,3d,1w,1M......
* @limit: 返回数据的条数限制(默认为1000，如果返回数据多于1000条，那么只返回1000条)
//当前上下文
ctx, _ := context.WithCancel(context.Background())
//
//创建一个新的服务
as := services.NewAPIService("Binance",
	ctx,
	log4.NewFileLogger("access-binance-spider.log"),
	log4.NewOutLogger())

res, err := as.Klines_Binance(info.BinanceKlinesRequestParams{
	Symbol:   "BTCUSDT",
	Interval: internal.Binance_Interval_Hour,
	Limit:    500,
})

if err != nil {
	fmt.Println(err)
} else {
	for _, v := range res {
		fmt.Println(v)
	}
}
*/
func (as *apiService) Klines_Binance(kr info.BinanceKlinesRequestParams) ([]*info.BinanceKline_Network, error) {
	params := make(map[string]string)
	params["symbol"] = kr.Symbol
	params["interval"] = string(kr.Interval)
	if kr.Limit != 0 {
		params["limit"] = strconv.Itoa(kr.Limit)
	}
	if kr.StartTime != 0 {
		params["startTime"] = strconv.FormatInt(kr.StartTime, 10)
	}
	if kr.EndTime != 0 {
		params["endTime"] = strconv.FormatInt(kr.EndTime, 10)
	}

	res, err := as.request_binance("GET", "/api/v1/klines", params, false, false)

	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from Klines")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.Wrap(err, "error unmarshal failed")
	}

	rawKlines := [][]interface{}{}
	if err := json.Unmarshal([]byte(textRes), &rawKlines); err != nil {
		return nil, errors.Wrap(err, "rawKlines unmarshal failed")
	}

	klines := []*info.BinanceKline_Network{}
	for _, k := range rawKlines {
		ot, err := timeFromUnixTimestampFloat(k[0])
		if err != nil {
			return nil, errors.Wrap(err, "cannot parse Kline.OpenTime")
		}
		open, err := floatFromString(k[1])
		if err != nil {
			return nil, errors.Wrap(err, "cannot parse Kline.Open")
		}
		high, err := floatFromString(k[2])
		if err != nil {
			return nil, errors.Wrap(err, "cannot parse Kline.High")
		}
		low, err := floatFromString(k[3])
		if err != nil {
			return nil, errors.Wrap(err, "cannot parse Kline.Low")
		}
		cls, err := floatFromString(k[4])
		if err != nil {
			return nil, errors.Wrap(err, "cannot parse Kline.Close")
		}
		volume, err := floatFromString(k[5])
		if err != nil {
			return nil, errors.Wrap(err, "cannot parse Kline.Volume")
		}
		ct, err := timeFromUnixTimestampFloat(k[6])
		if err != nil {
			return nil, errors.Wrap(err, "cannot parse Kline.CloseTime")
		}
		qav, err := floatFromString(k[7])
		if err != nil {
			return nil, errors.Wrap(err, "cannot parse Kline.QuoteAssetVolume")
		}
		not, ok := k[8].(float64)
		if !ok {
			return nil, errors.Wrap(err, "cannot parse Kline.NumberOfTrades")
		}
		tbbav, err := floatFromString(k[9])
		if err != nil {
			return nil, errors.Wrap(err, "cannot parse Kline.TakerBuyBaseAssetVolume")
		}
		tbqav, err := floatFromString(k[10])
		if err != nil {
			return nil, errors.Wrap(err, "cannot parse Kline.TakerBuyQuoteAssetVolume")
		}
		klines = append(klines, &info.BinanceKline_Network{
			OpenTime:                 ot,
			Open:                     open,
			High:                     high,
			Low:                      low,
			Close:                    cls,
			Volume:                   volume,
			CloseTime:                ct,
			QuoteAssetVolume:         qav,
			NumberOfTrades:           int(not),
			TakerBuyBaseAssetVolume:  tbbav,
			TakerBuyQuoteAssetVolume: tbqav,
		})
	}

	return klines, nil
}
