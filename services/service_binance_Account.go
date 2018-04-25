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
 * 获取 用户信息
 * @param  {[type]} as [description]
 * @return {[type]}    [description]
 //当前上下文
 ctx, _ := context.WithCancel(context.Background())

 //创建一个新的服务
 as := services.NewAPIService("Huobi",
	 ctx,
	 log4.NewFileLogger("aaa"),
	 log4.NewOutLogger())

 res, err := as.GetAccounts_Binance(info.BinanceAccountRequest_Network{
	 RecvWindow: 5 * time.Second,
	 Timestamp:  time.Now(),
 })
 if err != nil {
	 fmt.Printf("%v\n", err)
 } else {
	 fmt.Printf("%v\n", res)
 }
*/
func (as *apiService) GetAccounts_Binance(ar info.BinanceAccountRequest_Network) (*info.BinanceAccount_Network, error) {
	params := make(map[string]string)
	params["timestamp"] = strconv.FormatInt(unixMillis(ar.Timestamp), 10)
	if ar.RecvWindow != 0 {
		params["recvWindow"] = strconv.FormatInt(recvWindow(ar.RecvWindow), 10)
	}

	res, err := as.request_binance("GET", "/api/v3/account", params, true, true)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from account.get")
	}
	defer res.Body.Close()
	//
	// if res.StatusCode != 200 {
	// 	return nil, as.handleError(textRes)
	// }

	rawAccount := struct {
		MakerCommision   int64 `json:"makerCommision"`
		TakerCommission  int64 `json:"takerCommission"`
		BuyerCommission  int64 `json:"buyerCommission"`
		SellerCommission int64 `json:"sellerCommission"`
		CanTrade         bool  `json:"canTrade"`
		CanWithdraw      bool  `json:"canWithdraw"`
		CanDeposit       bool  `json:"canDeposit"`
		Balances         []struct {
			Asset  string `json:"asset"`
			Free   string `json:"free"`
			Locked string `json:"locked"`
		}
	}{}
	if err := json.Unmarshal(textRes, &rawAccount); err != nil {
		return nil, errors.Wrap(err, "rawAccount unmarshal failed")
	}

	acc := &info.BinanceAccount_Network{
		MakerCommision:  rawAccount.MakerCommision,
		TakerCommision:  rawAccount.TakerCommission,
		BuyerCommision:  rawAccount.BuyerCommission,
		SellerCommision: rawAccount.SellerCommission,
		CanTrade:        rawAccount.CanTrade,
		CanWithdraw:     rawAccount.CanWithdraw,
		CanDeposit:      rawAccount.CanDeposit,
	}
	for _, b := range rawAccount.Balances {
		f, err := floatFromString(b.Free)
		if err != nil {
			return nil, err
		}
		l, err := floatFromString(b.Locked)
		if err != nil {
			return nil, err
		}
		acc.Balances = append(acc.Balances, &info.BinanceBalance_Network{
			Asset:  b.Asset,
			Free:   f,
			Locked: l,
		})
	}

	return acc, nil
}

func (as *apiService) Ping_Binance() error {
	params := make(map[string]string)
	response, err := as.request_binance("GET", "/api/v1/ping", params, false, false)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", response.StatusCode)
	return nil
}
