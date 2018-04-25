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
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type apiFileLogger struct {
	LogFileWrite *logrus.Logger
	LogOutPut    *logrus.Logger
}

// New Creates a new logger with a "stderr" writer to send
// log messages at or above lvl to standard output.
func NewFileLogger(fileName string) Logger {
	var log = logrus.New()

	baseLogPath := fileName
	writer, _ := rotatelogs.New(
		baseLogPath+".%F",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)

	logFormat := &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true, DisableTimestamp: false, DisableColors: false, ForceColors: true, DisableSorting: true}
	// logFormat := &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05", DisableTimestamp: false}

	//参考:http://xiaorui.cc/2018/01/11/golang-logrus%E7%9A%84%E9%AB%98%E7%BA%A7%E9%85%8D%E7%BD%AEhook-logrotate/
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, logFormat)
	log.AddHook(lfHook)
	log.Formatter = &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true, DisableTimestamp: false, DisableColors: false, ForceColors: true}
	log.SetLevel(logrus.DebugLevel)
	log.Out = os.Stdout
	// log.SetFormatter(logFormat)

	return &apiFileLogger{
		LogFileWrite: log,
		LogOutPut:    nil,
	}
}

func (this *apiFileLogger) WithFields(fields logrus.Fields) *logrus.Entry {
	return this.LogFileWrite.WithFields(fields)
}

func (this *apiFileLogger) WithError(err error) *logrus.Entry {
	return this.LogOutPut.WithError(err)
}

func (this *apiFileLogger) Fatal(args ...interface{}) {
	this.LogFileWrite.Fatal(args)
}
func (this *apiFileLogger) Fatalf(format string, args ...interface{}) {
	this.LogFileWrite.Fatalf(format, args)
}

func (this *apiFileLogger) Debug(args ...interface{}) {
	this.LogFileWrite.Debug(args)
}
func (this *apiFileLogger) Debugf(format string, args ...interface{}) {
	this.LogFileWrite.Debugf(format, args)
}

func (this *apiFileLogger) Warning(args ...interface{}) {
	this.LogFileWrite.Warning(args)
}
func (this *apiFileLogger) Warningf(format string, args ...interface{}) {
	this.LogFileWrite.Warningf(format, args)
}

func (this *apiFileLogger) Info(args ...interface{}) {
	this.LogFileWrite.Info(args)
}
func (this *apiFileLogger) Infof(format string, args ...interface{}) {
	this.LogFileWrite.Infof(format, args)
}

func (this *apiFileLogger) Error(args ...interface{}) {
	this.LogFileWrite.Error(args)
}
func (this *apiFileLogger) Errorf(format string, args ...interface{}) {
	this.LogFileWrite.Errorf(format, args)
}
