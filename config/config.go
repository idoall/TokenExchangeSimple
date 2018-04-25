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

package config

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

// huobi config
type huobiConfig struct {
	OnLine_NewOrder bool //是否在线下单
	ACCESS_KEY      string
	SECRET_KEY      string
	MARKET_URL      string
	TRADE_URL       string
	SPOTID          int
}

// binance config
type binanceConfig struct {
	OnLine_NewOrder bool //是否在线下单
	API_KEY         string
	SECRET_KEY      string
	MARKET_URL      string
}

// zb config
type zbConfig struct {
	OnLine_NewOrder bool //是否在线下单
	ACCESS_KEY      string
	SECRET_KEY      string
	MARKET_URL      string
	TRADE_URL       string
}

type config struct {
	HuobiConfig   *huobiConfig
	BinanceConfig *binanceConfig
	ZBConfig      *zbConfig
}

var Config *config

/**
 * 初始化配置
 * @type {[type]}
 */
func Initialization() {
	fileName := "conf/my.ini"

	_c := &config{}

	//载入文件
	cfg, err := ini.Load(fileName)

	//读取huobi
	_huobiConfig := new(huobiConfig)
	err = cfg.Section("huobiConfig").MapTo(_huobiConfig)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	_c.HuobiConfig = _huobiConfig

	//读取 binance
	_binanceConfig := new(binanceConfig)
	err = cfg.Section("binanceConfig").MapTo(_binanceConfig)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	_c.BinanceConfig = _binanceConfig

	//读取 zb
	_zbConfig := new(zbConfig)
	err = cfg.Section("zbConfig").MapTo(_zbConfig)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	_c.ZBConfig = _zbConfig

	Config = _c
}
