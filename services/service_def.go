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
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/idoall/TokenExchangeSimple/config"
	"github.com/idoall/TokenExchangeSimple/info"
	"github.com/idoall/TokenExchangeSimple/log4"
	"github.com/idoall/TokenExchangeSimple/utils"
	"github.com/pkg/errors"
)

// Service represents service layer for Token Exchange API.
//
// The main purpose for this layer is to be replaced with dummy implementation
// if necessary without need to replace Binance instance.
type Service interface {
	GetAppName() string

	//----------------币安相关
	Ping_Binance() error

	//获取币安K线数据
	Klines_Binance(kr info.BinanceKlinesRequestParams) ([]*info.BinanceKline_Network, error)

	//获取用户信息
	GetAccounts_Binance(ar info.BinanceAccountRequest_Network) (*info.BinanceAccount_Network, error)

	//当前最新价格
	GetTradeNowPrice_Binance(strSymbol string) (float64, error)

	//Binance 委托下单
	NewOrder_Binance(placeRequestParams info.BinancePlaceRequestParams) (*info.BinancePlaceReturn_Network, error)

	//----------------ZB相关

	//获取用户信息
	GetAccounts_ZB() (*info.ZBAccountsReturn_Network, error)

	//获取ZB K线数据
	Klines_ZB(kr info.ZBKlinesRequestParams) (*info.ZBKLineReturn_Network, error)

	//获取  ZB 最新行情
	GetTicket_ZB(market string) (*info.ZBTicket_Network, error)

	//当前最新价格
	GetTradeNowPrice_ZB(strSymbol string) (float64, error)

	//ZB 委托下单
	NewOrder_ZB(placeRequestParams info.ZBPlaceRequestParams) (*info.ZBPlaceReturn_Network, error)

	//----------------火币相关

	//获取火币K线数据
	Klines_Huobi(kr info.HuobiKlinesRequestParams) (*info.HuobiKLineReturn_Network, error)

	//获取指定交易对的交易最新价格
	GetTradeNowPrice_Huobi(strSymbol string) (float64, error)

	//批量获取最近的交易记录
	GetTrade_Huobi(kr info.HuobiKlinesRequestParams) (*info.HuobiTradeReturn_Network, error)

	//获取用户信息
	GetAccounts_Huobi()

	//下订单
	NewOrder_Huobi(placeRequestParams info.HuobiPlaceRequestParams) (*info.HuobiPlaceReturn_Network, error)
}

/**
 * API服务
 * @type {[type]}
 */
type apiService struct {
	AppName        string
	Ctx            context.Context
	Log4FileWriter log4.Logger
	Log4OutPut     log4.Logger
}

// NewAPIService creates instance of Service.
//
func NewAPIService(appName string, ctx context.Context, _log4File, _log4Out log4.Logger) Service {

	return &apiService{
		AppName:        appName,
		Ctx:            ctx,
		Log4FileWriter: _log4File,
		Log4OutPut:     _log4Out,
	}
}

func (as *apiService) GetAppName() string {
	return as.AppName
}

func (as *apiService) request_binance(method string, endpoint string, params map[string]string,
	apiKey bool, sign bool) (*http.Response, error) {
	transport := &http.Transport{}
	client := &http.Client{
		Transport: transport,
	}

	url := config.Config.BinanceConfig.MARKET_URL + endpoint

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create request")
	}
	req.WithContext(as.Ctx)

	q := req.URL.Query()
	sortedParams := utils.MapSortByKey(params)
	for key, val := range sortedParams {
		q.Add(key, val)
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")

	if apiKey {
		req.Header.Add("X-MBX-APIKEY", config.Config.BinanceConfig.API_KEY)
	}
	if sign {
		hmacSigner := &BinanceHmacSigner{
			Key: []byte(config.Config.BinanceConfig.SECRET_KEY),
		}

		q.Add("signature", hmacSigner.Sign([]byte(q.Encode())))
	}
	req.URL.RawQuery = q.Encode()
	// fmt.Println(req.URL)
	resp, err := client.Do(req)
	// fmt.Println(resp)
	// fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/**
 * 火币的 request 请求
 * @param  {[type]} as [description]
 * @return {[type]}    [description]
 */
func (as *apiService) request_Huobi(method string, strRequestPath string, params map[string]string, sign bool) (*http.Response, error) {
	//生成client 参数为默认
	transport := &http.Transport{}
	client := &http.Client{
		Transport: transport,
	}

	url := config.Config.HuobiConfig.MARKET_URL + strRequestPath

	//火币的GET方法，对于GET请求，每个方法自带的参数都需要进行签名运算
	if method == "GET" {

		//返回一个新的服务器访问请求， 这个请求可以传递给 http.Handler 以便进行测试
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "unable to create request")
		}

		req.WithContext(as.Ctx)
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")

		var strParams string
		//如果需要每签名
		if sign {
			//对传入的字符串进行ASCII排序和拼接
			sortedParams := utils.MapSortByKey(params)
			encodeParams := utils.MapValueEncodeURI(sortedParams)
			strParams = utils.Map2UrlQuery(encodeParams)

			//使用 HMAC 进行签名计算
			hmacSigner := &HuobiHmacSigner{
				Key: []byte(config.Config.HuobiConfig.SECRET_KEY),
			}
			hostName := "api.huobi.pro"
			strPayload := method + "\n" + hostName + "\n" + strRequestPath + "\n" + strParams
			params["Signature"] = hmacSigner.Sign([]byte(strPayload))

			//按照ASCII码的顺序对参数名进行排序
			strParams = utils.Map2UrlQuery(params)
		} else {
			//按照ASCII码的顺序对参数名进行排序
			strParams = utils.Map2UrlQuery(params)
		}

		req.URL.RawQuery = strParams
		//根据客户端配置的策略（例如重定向，Cookie，身份验证）发送http请求并返回http响应
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	} else if method == "POST" {
		//如果是POST方法，对于POST请求，每个方法自带的参数不进行签名认证，即POST请求中需要进行签名运算的只有AccessKeyId、SignatureMethod、SignatureVersion、Timestamp四个参数，其它参数放在body中

		//将请求的参数放入Body中
		jsonParams := ""
		if nil != params {
			bytesParams, _ := json.Marshal(params)
			jsonParams = string(bytesParams)
		}
		req, err := http.NewRequest(method, url, strings.NewReader(jsonParams))
		if err != nil {
			return nil, errors.Wrap(err, "unable to create request")
		}

		req.WithContext(as.Ctx)
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept-Language", "zh-cn")

		//如果需要签名
		var strParams string
		if sign {
			//配置要签名的参数
			signmapParams := make(map[string]string)
			signmapParams["AccessKeyId"] = config.Config.HuobiConfig.ACCESS_KEY
			signmapParams["SignatureMethod"] = "HmacSHA256"
			signmapParams["SignatureVersion"] = "2"
			signmapParams["Timestamp"] = time.Now().UTC().Format("2006-01-02T15:04:05")

			//对参数的字符串进行ASCII排序和拼接
			signsortedParams := utils.MapSortByKey(signmapParams)
			signencodeParams := utils.MapValueEncodeURI(signsortedParams)
			signstrParams := utils.Map2UrlQuery(signencodeParams)

			//使用Hmac进行签名
			hmacSigner := &HuobiHmacSigner{
				Key: []byte(config.Config.HuobiConfig.SECRET_KEY),
			}
			hostName := "api.huobi.pro"
			strPayload := method + "\n" + hostName + "\n" + strRequestPath + "\n" + signstrParams
			// fmt.Println("params", params)
			signmapParams["Signature"] = hmacSigner.Sign([]byte(strPayload))

			//对签名参数排序
			strParams = utils.Map2UrlQuery(signmapParams)
			// fmt.Println("strParams", strParams)
		}
		req.URL.RawQuery = strParams
		// fmt.Println("req.URL", req.URL)
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	return nil, errors.New("错误的 Method 参数")
}

func (as *apiService) request_ZB(method string, strRequestPath string, params map[string]string, sign bool) (*http.Response, error) {
	transport := &http.Transport{}
	client := &http.Client{
		Transport: transport,
	}

	//添加必要参数accesskey
	params["accesskey"] = config.Config.ZBConfig.ACCESS_KEY

	// url := fmt.Sprintf("%s/%s", config.Config.BinanceConfig.MARKET_URL, endpoint)

	req, err := http.NewRequest(method, strRequestPath, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create request")
	}
	req.WithContext(as.Ctx)

	q := req.URL.Query()
	sortedParams := utils.MapSortByKey(params)
	for key, val := range sortedParams {
		q.Add(key, val)
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")

	var strParams string
	if sign {
		if params["method"] == "order" {
			//先用sha加密secretkey
			hmacSigner := &ZBHmacSigner{
				Key: []byte(config.Config.ZBConfig.SECRET_KEY),
			}

			//对参数的字符串进行ASCII排序和拼接
			signsortedParams := utils.MapSortByKey(params)
			signencodeParams := utils.MapValueEncodeURI(signsortedParams)
			signstrParams := utils.Map2UrlQuery(signencodeParams)
			params["sign"] = hmacSigner.Sign([]byte(signstrParams))
			params["reqTime"] = fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

			//对签名参数排序
			signencodeParams = utils.MapValueEncodeURI(params)
			strParams = utils.Map2UrlQuery(signencodeParams)
		} else {

			//先用sha加密secretkey
			hmacSigner := &ZBHmacSigner{
				Key: []byte(config.Config.ZBConfig.SECRET_KEY),
			}

			//对参数的字符串进行ASCII排序和拼接
			signsortedParams := utils.MapSortByKey(params)
			signencodeParams := utils.MapValueEncodeURI(signsortedParams)
			signstrParams := utils.Map2UrlQuery(signencodeParams)

			mapParams2Sign := make(map[string]string)
			mapParams2Sign["accesskey"] = config.Config.ZBConfig.ACCESS_KEY
			mapParams2Sign["method"] = params["method"]
			mapParams2Sign["sign"] = hmacSigner.Sign([]byte(signstrParams))
			mapParams2Sign["reqTime"] = fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

			//对签名参数排序
			signencodeParams = utils.MapValueEncodeURI(mapParams2Sign)
			strParams = utils.Map2UrlQuery(signencodeParams)
		}
	} else {
		strParams = utils.Map2UrlQuery(params)
	}
	req.URL.RawQuery = strParams
	// fmt.Println(req.URL)
	resp, err := client.Do(req)
	// fmt.Println(resp)
	// fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func floatFromString(raw interface{}) (float64, error) {
	str, ok := raw.(string)
	if !ok {
		return 0, errors.New(fmt.Sprintf("unable to parse, value not string: %T", raw))
	}
	flt, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("unable to parse, value not string: %T", raw))
	}
	return flt, nil
}

func intFromString(raw interface{}) (int, error) {
	str, ok := raw.(string)
	if !ok {
		return 0, errors.New(fmt.Sprintf("unable to parse, value not string: %T", raw))
	}
	n, err := strconv.Atoi(str)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("unable to parse as int: %T", raw))
	}
	return n, nil
}

func timeFromUnixTimestampFloat(raw interface{}) (time.Time, error) {
	ts, ok := raw.(float64)
	if !ok {
		return time.Time{}, errors.New(fmt.Sprintf("unable to parse, value not int64: %T", raw))
	}
	return time.Unix(0, int64(ts)*int64(time.Millisecond)), nil
}

func unixMillis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func recvWindow(d time.Duration) int64 {
	return int64(d) / int64(time.Millisecond)
}
