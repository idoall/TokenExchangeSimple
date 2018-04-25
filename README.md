数字货币交易所获取API数据相关实例:

* 支持查看帐号，查看K线，下订单

* 暂时只支持火币、币安、ZB

* 使用前请先配置conf/my.ini中的平台秘钥

Linux:
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

Mac:
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
