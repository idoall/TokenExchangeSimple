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
 * 获取当前交易对的最新价格
 * @param  {[type]} as [description]
 * @return {[type]}    [description]
 */
func (as *apiService) GetTradeNowPrice_Huobi(strSymbol string) (float64, error) {
	//获取最后一条交易记录的信息
	_model, err := as.GetTrade_Huobi(info.HuobiKlinesRequestParams{
		Symbol: strSymbol,
		Size:   1,
	})

	if err != nil {
		return 0, errors.Wrap(err, "GetTradeNowPrice_Huobi")
	}

	if _model.Status == "ok" {
		return _model.Data[0].Data[0].Price, nil
	}

	return 0, errors.Errorf("[GetTradeNowPrice_Huobi]Symbol:%s	err-code:%s	err-msg:%s", strSymbol, _model.ErrCode, _model.ErrMsg)

}

/**
* 批量获取最近的交易记录
* @param  {[type]} as [description]
* @return {[type]}    [description]
ctx, _ := context.WithCancel(context.Background())
as := services.NewAPIService("aa", ctx)
res, err := as.Klines_Huobi(info.HuobiKlinesRequestParams{
	Symbol: "btcusdt",
	Period: "5min",
	Size:   100,
})
if err != nil {
	fmt.Println(err)
}
fmt.Println(res)
*/
func (as *apiService) GetTrade_Huobi(kr info.HuobiKlinesRequestParams) (*info.HuobiTradeReturn_Network, error) {
	//返回对象
	tradeReturn := info.HuobiTradeReturn_Network{}

	//请求参数
	mapParams := make(map[string]string)
	mapParams["symbol"] = kr.Symbol
	mapParams["size"] = strconv.Itoa(kr.Size)

	//发送请求
	res, err := as.request_Huobi("GET", "/market/history/trade", mapParams, false)
	if err != nil {
		return &info.HuobiTradeReturn_Network{}, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &info.HuobiTradeReturn_Network{}, errors.Wrap(err, "unable to read response from info.HuobiTradeReturn_Network")
	}
	defer res.Body.Close()

	//json解析
	json.Unmarshal([]byte(textRes), &tradeReturn)
	// fmt.Println("symbol", kr.Symbol)
	// fmt.Printf("%v", tradeReturn)

	return &tradeReturn, nil
}
