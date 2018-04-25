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
 * 获取  ZB 最新行情
 * @param  {[type]} as [description]
 * @return {[type]}    [description]
 */
func (as *apiService) GetTicket_ZB(market string) (*info.ZBTicket_Network, error) {

	strRequest := "/data/v1/ticker"

	mapParams := make(map[string]string)
	mapParams["market"] = market

	strRequestUrl := config.Config.ZBConfig.MARKET_URL + strRequest

	res, err := as.request_ZB("GET", strRequestUrl, mapParams, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from Ticket.get")
	}
	defer res.Body.Close()
	// fmt.Println("textRes", string(textRes))
	_m := info.ZBTicket_Network{}
	if err := json.Unmarshal([]byte(textRes), &_m); err != nil {
		return nil, errors.Wrap(err, "ZBTicket_Network unmarshal failed")
	}

	return &_m, nil
}
