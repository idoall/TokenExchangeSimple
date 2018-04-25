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
	"github.com/idoall/TokenExchangeSimple/info"
)

// Ema is the main object
type KDJ_ZB struct {
	Period int //默认计算几天的KDJ一般都是9天的数据
	RSV    []float64
	K      []float64
	D      []float64
	J      []float64
}

func NewKdj_ZB(period int) *KDJ_ZB {
	_kdj := &KDJ_ZB{Period: 9, RSV: make([]float64, 0), K: make([]float64, 0), D: make([]float64, 0), J: make([]float64, 0)}
	return _kdj
}

/*
计算KDJ

RSV = (CLOSE-LLV(LOW,9))/(HHV(HIGH,9)-LLV(LOW,9))*100;
K = SMA(RSV,3,1);
D = SMA(K,3,1);
J = 3*K-2*D;

LLV means lowest value in period
HHV means highest value in period
SMA is Simple Moving Average
*/
func (_kdj *KDJ_ZB) CalculationKDJ(records []info.ZBKline_DB) {

	_kdj.calculationKD(records)
	_kdj.CalculationJ()

	_kdj.K = ReverseFloat64(_kdj.K)
	_kdj.D = ReverseFloat64(_kdj.D)
	_kdj.J = ReverseFloat64(_kdj.J)
	_kdj.RSV = ReverseFloat64(_kdj.RSV)
}

/**
 * 计算出kd值
 * @param  {[type]} _kdj [description]
 * @return {[type]}      [description]
 */
func (_kdj *KDJ_ZB) calculationKD(records []info.ZBKline_DB) {

	var periodLowArr, periodHighArr []float64
	length := len(records)
	var rsv []float64 = make([]float64, length)
	var k []float64 = make([]float64, length)
	var d []float64 = make([]float64, length)

	// Loop through the entire array.
	for i := 0; i < length; i++ {
		// add points to the array.
		periodLowArr = append(periodLowArr, records[i].Low)
		periodHighArr = append(periodHighArr, records[i].High)

		// 1: Check if array is "filled" else create null point in line.
		// 2: Calculate average.
		// 3: Remove first value.

		if _kdj.Period == len(periodLowArr) {
			lowest := _kdj.arrayLowest(periodLowArr)
			highest := _kdj.arrayHighest(periodHighArr)
			//logger.Infoln(i, records[i].Close, lowest, highest)
			if highest-lowest < 0.000001 {
				rsv[i] = 100
			} else {
				rsv[i] = (records[i].Close - lowest) / (highest - lowest) * 100
			}

			// k[i] = (rsv[i] + 2.0*k[i-1]) / 3
			// d[i] = (k[i] + 2.0*d[i-1]) / 3
			k[i] = (2.0/3)*k[i-1] + 1.0/3*rsv[i]
			d[i] = (2.0/3)*d[i-1] + 1.0/3*k[i]
			// remove first value in array.
			periodLowArr = periodLowArr[1:]
			periodHighArr = periodHighArr[1:]
		} else {
			k[i] = 50
			d[i] = 50
			rsv[i] = 0
		}
	}

	_kdj.RSV = rsv
	_kdj.K = k
	_kdj.D = d
}

/**
 * 计算J值
 * @param  {[type]} _kdj [description]
 * @return {[type]}      [description]
 */
func (_kdj *KDJ_ZB) CalculationJ() {
	length := len(_kdj.K)
	var j []float64 = make([]float64, length)

	// Loop through the entire array.
	for i := 0; i < length; i++ {
		j[i] = 3*_kdj.K[i] - 2*_kdj.D[i]

	}
	_kdj.J = j
}

func (_kdj *KDJ_ZB) highest(Price []float64, periods int) []float64 {
	var periodArr []float64
	length := len(Price)
	var HighestLine []float64 = make([]float64, length)

	// Loop through the entire array.
	for i := 0; i < length; i++ {
		// add points to the array.
		periodArr = append(periodArr, Price[i])
		// 1: Check if array is "filled" else create null point in line.
		// 2: Calculate average.
		// 3: Remove first value.
		if periods == len(periodArr) {
			HighestLine[i] = _kdj.arrayHighest(periodArr)

			// remove first value in array.
			periodArr = periodArr[1:]
		} else {
			HighestLine[i] = 0
		}
	}

	return HighestLine
}

func (_kdj *KDJ_ZB) lowest(Price []float64, periods int) []float64 {
	var periodArr []float64
	length := len(Price)
	var LowestLine []float64 = make([]float64, length)

	// Loop through the entire array.
	for i := 0; i < length; i++ {
		// add points to the array.
		periodArr = append(periodArr, Price[i])
		// 1: Check if array is "filled" else create null point in line.
		// 2: Calculate average.
		// 3: Remove first value.
		if periods == len(periodArr) {
			LowestLine[i] = _kdj.arrayLowest(periodArr)

			// remove first value in array.
			periodArr = periodArr[1:]
		} else {
			LowestLine[i] = 0
		}
	}

	return LowestLine
}

func (_kdj *KDJ_ZB) arrayLowest(Price []float64) float64 {
	length := len(Price)
	var lowest = Price[0]

	// Loop through the entire array.
	for i := 1; i < length; i++ {
		if Price[i] < lowest {
			lowest = Price[i]
		}
	}

	return lowest
}

func (_kdj *KDJ_ZB) arrayHighest(Price []float64) float64 {
	length := len(Price)
	var highest = Price[0]

	// Loop through the entire array.
	for i := 1; i < length; i++ {
		if Price[i] > highest {
			highest = Price[i]
		}
	}

	return highest
}

// func Highest(Price []float64, periods int) []float64 {
// 	var periodArr []float64
// 	length := len(Price)
// 	var HighestLine []float64 = make([]float64, length)
//
// 	// Loop through the entire array.
// 	for i := 0; i < length; i++ {
// 		// add points to the array.
// 		periodArr = append(periodArr, Price[i])
// 		// 1: Check if array is "filled" else create null point in line.
// 		// 2: Calculate average.
// 		// 3: Remove first value.
// 		if periods == len(periodArr) {
// 			HighestLine[i] = arrayHighest(periodArr)
//
// 			// remove first value in array.
// 			periodArr = periodArr[1:]
// 		} else {
// 			HighestLine[i] = 0
// 		}
// 	}
//
// 	return HighestLine
// }

// func Lowest(Price []float64, periods int) []float64 {
// 	var periodArr []float64
// 	length := len(Price)
// 	var LowestLine []float64 = make([]float64, length)
//
// 	// Loop through the entire array.
// 	for i := 0; i < length; i++ {
// 		// add points to the array.
// 		periodArr = append(periodArr, Price[i])
// 		// 1: Check if array is "filled" else create null point in line.
// 		// 2: Calculate average.
// 		// 3: Remove first value.
// 		if periods == len(periodArr) {
// 			LowestLine[i] = arrayLowest(periodArr)
//
// 			// remove first value in array.
// 			periodArr = periodArr[1:]
// 		} else {
// 			LowestLine[i] = 0
// 		}
// 	}
//
// 	return LowestLine
// }

// func arrayLowest(Price []float64) float64 {
// 	length := len(Price)
// 	var lowest = Price[0]
//
// 	// Loop through the entire array.
// 	for i := 1; i < length; i++ {
// 		if Price[i] < lowest {
// 			lowest = Price[i]
// 		}
// 	}
//
// 	return lowest
// }

// func arrayHighest(Price []float64) float64 {
// 	length := len(Price)
// 	var highest = Price[0]
//
// 	// Loop through the entire array.
// 	for i := 1; i < length; i++ {
// 		if Price[i] > highest {
// 			highest = Price[i]
// 		}
// 	}
//
// 	return highest
// }
