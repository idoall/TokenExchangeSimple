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

// Ema is the main object
type Ema struct {
	period int
	points []point
}

type point struct {
	Timestamp int64
	Value     float64 //数值
	Ema       float64 //计算后的EMA值
}

func NewEma(period int) *Ema {
	ema := &Ema{period: period, points: make([]point, 0)}
	return ema
}

func (ema *Ema) GetPoints() []point {
	return ema.points
}

// Add adds a new Value to Ema
func (ema *Ema) Add(timestamp int64, value float64) {
	p := point{Timestamp: timestamp, Value: value}

	//平滑指数，一般取作2/(N+1)
	alpha := 2.0 / (float64(ema.period) + 1.0)

	// fmt.Println(alpha)

	emaTminusOne := value
	if len(ema.points) > 0 {
		emaTminusOne = ema.points[len(ema.points)-1].Ema
	}

	emaT := alpha*value + (1-alpha)*emaTminusOne
	p.Ema = emaT
	ema.points = append(ema.points, p)
}
