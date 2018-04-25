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

package log4

import (
	"github.com/sirupsen/logrus"
)

// Version information
const (
	VERSION = "v0.1.0"
	MAJOR   = 0
	MINOR   = 1
	BUILD   = 0
)

type Logger interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	// Fine(arg0 interface{}, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	// Trace(arg0 interface{}, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	WithFields(fields logrus.Fields) *logrus.Entry

	WithError(err error) *logrus.Entry
}
