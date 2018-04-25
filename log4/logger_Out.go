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
	"os"

	"github.com/sirupsen/logrus"
)

type apiOutLogger struct {
	LogFileWrite *logrus.Logger
	LogOutPut    *logrus.Logger
}

// New Creates a new logger with a "stderr" writer to send
// log messages at or above lvl to standard output.
func NewOutLogger() Logger {
	var log = logrus.New()

	log.Formatter = &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true, DisableTimestamp: false, DisableColors: false, ForceColors: true, DisableSorting: true}
	log.SetLevel(logrus.DebugLevel)
	log.Out = os.Stdout
	// log.SetFormatter(logFormat)

	return &apiOutLogger{
		LogFileWrite: nil,
		LogOutPut:    log,
	}
}

func (this *apiOutLogger) WithFields(fields logrus.Fields) *logrus.Entry {
	return this.LogOutPut.WithFields(fields)
}

func (this *apiOutLogger) WithError(err error) *logrus.Entry {
	return this.LogOutPut.WithError(err)
}

func (this *apiOutLogger) Fatal(args ...interface{}) {
	this.LogOutPut.Fatal(args)
}
func (this *apiOutLogger) Fatalf(format string, args ...interface{}) {
	this.LogFileWrite.Fatalf(format, args)
}

func (this *apiOutLogger) Debug(args ...interface{}) {
	this.LogOutPut.Debug(args)
}
func (this *apiOutLogger) Debugf(format string, args ...interface{}) {
	this.LogOutPut.Debugf(format, args)
}

func (this *apiOutLogger) Warning(args ...interface{}) {
	this.LogOutPut.Warning(args)
}
func (this *apiOutLogger) Warningf(format string, args ...interface{}) {
	this.LogOutPut.Warningf(format, args)
}

func (this *apiOutLogger) Info(args ...interface{}) {
	this.LogOutPut.Info(args)
}
func (this *apiOutLogger) Infof(format string, args ...interface{}) {
	this.LogOutPut.Infof(format, args)
}

func (this *apiOutLogger) Error(args ...interface{}) {
	this.LogOutPut.Error(args)
}
func (this *apiOutLogger) Errorf(format string, args ...interface{}) {
	this.LogOutPut.Errorf(format, args)
}
