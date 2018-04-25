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

package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"hash/crc32"
	"io"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/idoall/TokenExchangeSimple/info"
	"golang.org/x/net/html/charset"
)

/**
 * Binance手续费
 * @type {[type]}
 */
func ServiceCharge_Binance(amount float64) float64 {
	return amount - amount*0.001
}

/**
 * Huobi手续费
 * @type {[type]}
 */
func ServiceCharge_Huobi(amount float64) float64 {
	return amount - amount*0.002
}

/**
 * ZB手续费
 * @type {[type]}
 */
func ServiceCharge_ZB(amount float64) float64 {
	return amount - amount*0.002
}

/**
 * 获取调用方法的名称
 * @type {[type]}
 */
func GetFuncName() string {
	// Skip this function, and fetch the PC and file for its parent
	pc, _, _, _ := runtime.Caller(1)
	// Retrieve a Function object this functions parent
	functionObject := runtime.FuncForPC(pc)
	// Regex to extract just the function name (and not the module path)
	extractFnName := regexp.MustCompile(`^.*\.(.*)$`)
	return extractFnName.ReplaceAllString(functionObject.Name(), "$1")
}

/**
 * 获取调用方法的行号
 * @type {[type]}
 */
func GetFuncLine() int {
	// Skip this function, and fetch the PC and file for its parent
	_, _, line, _ := runtime.Caller(1)
	return line
}

func HmacMD5(key, data string) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	return hex.EncodeToString(hmac.Sum([]byte("")))
}

func Sha1ToHex(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// 对Map按着ASCII码进行排序
// mapValue: 需要进行排序的map
// return: 排序后的map
func MapSortByKey(mapValue map[string]string) map[string]string {
	var keys []string
	for key := range mapValue {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	mapReturn := make(map[string]string)
	for _, key := range keys {
		mapReturn[key] = mapValue[key]
	}

	return mapReturn
}

// 将map格式的请求参数转换为字符串格式的
// mapParams: map格式的参数键值对
// return: 查询字符串
func Map2UrlQuery(mapParams map[string]string) string {

	//golang 的 range 有随机性，确保稳定性
	sorted_keys := make([]string, 0)
	for k, _ := range mapParams {
		sorted_keys = append(sorted_keys, k)
	}
	// sort 'string' key in increasing order
	sort.Strings(sorted_keys)

	var strParams string
	for _, key := range sorted_keys {
		strParams += (key + "=" + mapParams[key] + "&")
	}

	if 0 < len(strParams) {
		strParams = string([]rune(strParams)[:len(strParams)-1])
	}

	return strParams
}

// 对Map的值进行URI编码
// mapParams: 需要进行URI编码的map
// return: 编码后的map
func MapValueEncodeURI(mapValue map[string]string) map[string]string {
	for key, value := range mapValue {
		valueEncodeURI := url.QueryEscape(value)
		mapValue[key] = valueEncodeURI
	}

	return mapValue
}

/**
 * 对数组进行反转到一个新数组
 * @type {[type]}
 */
func ReverseLineData_Binance(s []info.BinanceKline_DB) []info.BinanceKline_DB {
	_length := len(s)
	var _tempArr []info.BinanceKline_DB

	for i := _length - 1; i >= 0; i-- {
		_tempArr = append(_tempArr, s[i])
	}
	return _tempArr
}

/**
 * 对数组进行反转到一个新数组
 * @type {[type]}
 */
func ReverseLineData_Huobi(s []info.HuobiKline_DB) []info.HuobiKline_DB {
	_length := len(s)
	var _tempArr []info.HuobiKline_DB

	for i := _length - 1; i >= 0; i-- {
		_tempArr = append(_tempArr, s[i])
	}
	return _tempArr
}

/**
 * 对数组进行反转到一个新数组
 * @type {[type]}
 */
func ReverseLineData_ZB(s []info.ZBKline_DB) []info.ZBKline_DB {
	_length := len(s)
	var _tempArr []info.ZBKline_DB

	for i := _length - 1; i >= 0; i-- {
		_tempArr = append(_tempArr, s[i])
	}
	return _tempArr
}

/**
 * 对数组进行反转返回一个新数组
 * @type {[type]}
 */
func ReverseFloat64(s []float64) []float64 {
	_length := len(s)
	var _tempArr []float64

	for i := _length - 1; i >= 0; i-- {
		_tempArr = append(_tempArr, s[i])
	}
	return _tempArr
}

/**
 * 将字符串替换完以后，转成float并返回值
 * Example: replacesSpaceFloat("324,324.34 CNy", "CNY", ",")
 */
func ReplacesSpaceFloat(_str string, _oldString ...string) float64 {
	tempstr := _str

	//将后面的字符串替换成空
	for i := 0; i < len(_oldString); i++ {
		tempstr = strings.Replace(tempstr, _oldString[i], "", -1)
	}

	//去除空格
	tempstr = strings.TrimSpace(tempstr)

	// tempstr = strings.Replace(tempstr, "\n", "", -1)
	//转换成float
	_float, err := strconv.ParseFloat(tempstr, 64)

	if err != nil {
		println(err)
		return 0
	} else {
		return _float
	}

}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// JsonpToJson modify jsonp string to json string
// Example: forbar({a:"1",b:2}) to {"a":"1","b":2}
func JsonpToJson(json string) string {
	start := strings.Index(json, "{")
	end := strings.LastIndex(json, "}")
	start1 := strings.Index(json, "[")
	if start1 > 0 && start > start1 {
		start = start1
		end = strings.LastIndex(json, "]")
	}
	if end > start && end != -1 && start != -1 {
		json = json[start : end+1]
	}
	json = strings.Replace(json, "\\'", "", -1)
	regDetail, _ := regexp.Compile("([^\\s\\:\\{\\,\\d\"]+|[a-z][a-z\\d]*)\\s*\\:")
	return regDetail.ReplaceAllString(json, "\"$1\":")
}

// The GetWDPath gets the work directory path.
func GetWDPath() string {
	wd := os.Getenv("GOPATH")
	if wd == "" {
		panic("GOPATH is not setted in env.")
	}
	return wd
}

// The IsDirExists judges path is directory or not.
func IsDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}

	panic("util isDirExists not reached")
}

// The IsFileExists judges path is file or not.
func IsFileExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return !fi.IsDir()
	}

	panic("util isFileExists not reached")
}

// The IsNum judges string is number or not.
func IsNum(a string) bool {
	reg, _ := regexp.Compile("^\\d+$")
	return reg.MatchString(a)
}

// simple xml to string  support utf8
func XML2mapstr(xmldoc string) map[string]string {
	var t xml.Token
	var err error
	inputReader := strings.NewReader(xmldoc)
	decoder := xml.NewDecoder(inputReader)
	decoder.CharsetReader = func(s string, r io.Reader) (io.Reader, error) {
		return charset.NewReader(r, s)
	}
	m := make(map[string]string, 32)
	key := ""
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			key = token.Name.Local
		case xml.CharData:
			content := string([]byte(token))
			m[key] = content
		default:
			// ...
		}
	}

	return m
}

//string to hash
func MakeHash(s string) string {
	const IEEE = 0xedb88320
	var IEEETable = crc32.MakeTable(IEEE)
	hash := fmt.Sprintf("%x", crc32.Checksum([]byte(s), IEEETable))
	return hash
}
