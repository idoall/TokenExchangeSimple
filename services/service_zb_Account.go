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

	"github.com/idoall/TokenExchangeSimple/config"
	"github.com/idoall/TokenExchangeSimple/info"
	"github.com/pkg/errors"
)

/**
 * 获取ZB用户信息
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

 //查询 ZB 用户信息
 res, err := as.GetAccounts_ZB()
 if err != nil {
	 fmt.Println(err)
 } else {
	 fmt.Printf("%v", res)
 }
*/
func (as *apiService) GetAccounts_ZB() (*info.ZBAccountsReturn_Network, error) {
	strRequest := "/api/getAccountInfo"
	mapParams := make(map[string]string)
	mapParams["method"] = "getAccountInfo"

	strRequestUrl := config.Config.ZBConfig.TRADE_URL + strRequest

	res, err := as.request_ZB("GET", strRequestUrl, mapParams, true)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from account.get")
	}
	defer res.Body.Close()

	rawAccount := struct {
		Result struct {
			Coins []struct {
				Freez       string `json:"freez"`
				EnName      string `json:"enName"`
				UnitDecimal int    `json:"unitDecimal"`
				UnName      string `json:"cnName"`
				UnitTag     string `json:"unitTag"`
				Available   string `json:"available"`
				Key         string `json:"key"`
			}
			Base struct {
				UserName               string `json:"username"`
				Trade_Password_Enabled bool   `json:"trade_password_enabled"`
				Auth_Google_Enabled    bool   `json:"auth_google_enabled"`
				Auth_Mobile_Enabled    bool   `json:"auth_mobile_enabled"`
			}
		}
	}{}
	if err := json.Unmarshal(textRes, &rawAccount); err != nil {
		return nil, errors.Wrap(err, "rawAccount unmarshal failed")
	}

	acc := &info.ZBAccountsReturn_Network{
		UserName:               rawAccount.Result.Base.UserName,
		Trade_Password_Enabled: rawAccount.Result.Base.Trade_Password_Enabled,
		Auth_Google_Enabled:    rawAccount.Result.Base.Auth_Google_Enabled,
		Auth_Mobile_Enabled:    rawAccount.Result.Base.Auth_Mobile_Enabled,
	}

	for _, b := range rawAccount.Result.Coins {
		f, err := floatFromString(b.Freez)
		if err != nil {
			return nil, err
		}
		a, err := floatFromString(b.Available)
		if err != nil {
			return nil, err
		}
		acc.Coins = append(acc.Coins, &info.ZBAccountsReturn_Coins_Network{
			Freez:       f,
			EnName:      b.EnName,
			UnitDecimal: b.UnitDecimal,
			UnName:      b.UnName,
			UnitTag:     b.UnitTag,
			Available:   a,
			Key:         b.Key,
		})
	}

	return acc, nil

}
