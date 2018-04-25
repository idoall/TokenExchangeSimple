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
	"strconv"

	"github.com/pkg/errors"
)

/**
 * 获取指定交易对的交易最新价格
 */
func (as *apiService) GetTradeNowPrice_ZB(strSymbol string) (float64, error) {
	//获取最后一条交易记录的信息
	_model, err := as.GetTicket_ZB(strSymbol)

	if err != nil {
		return 0, errors.Wrap(err, "GetTradeNowPrice_ZB->GetTicket_ZB")
	}

	// fmt.Println("_model", _model)
	price, err := strconv.ParseFloat(_model.Ticket.Last, 64)
	if err != nil {
		return 0, errors.Wrap(err, "GetTradeNowPrice_ZB ParseFloat")
	}

	return price, nil
}
