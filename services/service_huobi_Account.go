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
	"time"

	"github.com/idoall/TokenExchangeSimple/config"
	"github.com/idoall/TokenExchangeSimple/info"
)

/**
 * 获取帐号信息
 * @param  {[type]} as [description]
 * @return {[type]}    [description]
 //当前上下文
 ctx, _ := context.WithCancel(context.Background())

 //创建一个新的服务
 as := services.NewAPIService("Huobi",
	 config.InitConfig("conf/my.ini"),
	 ctx,
	 log4.NewFileLogger("aaa"),
	 log4.NewOutLogger())

 as.GetAccounts_Huobi()
*/
func (as *apiService) GetAccounts_Huobi() {
	strRequest := "/v1/account/accounts"
	strMethod := "GET"
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05")

	mapParams := make(map[string]string)
	mapParams["AccessKeyId"] = config.Config.HuobiConfig.ACCESS_KEY
	mapParams["SignatureMethod"] = "HmacSHA256"
	mapParams["SignatureVersion"] = "2"
	mapParams["Timestamp"] = timestamp

	res, err := as.request_Huobi(strMethod, strRequest, mapParams, true)
	if err != nil {
		fmt.Println(err)
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	accountsReturn := info.HuobiAccountsReturn_Network{}
	json.Unmarshal([]byte(textRes), &accountsReturn)

	for i := 0; i < len(accountsReturn.Data); i++ {
		var _item = accountsReturn.Data[i]
		switch _item.Type {
		case "spot":
			fmt.Printf("现货帐户：%s	%d	状态：%s	\n", _item.Type, _item.ID, _item.State)
		case "otc":
			fmt.Printf("otc帐户：%s	%d	状态：%s	\n", _item.Type, _item.ID, _item.State)
		default:
			fmt.Printf("未知帐户：%s	%d	状态：%s	\n", _item.Type, _item.ID, _item.State)
		}
		//fmt.Println(_accounts.Data[i])
	}
	// strUrl := config.Huobi_TRADE_URL + strRequestPath
	//
	// jsonAccountsReturn, err := NewHuobiFunc().Huobi_ApiKeyGet(make(map[string]string), strRequest)
}
