package main

import (
	"github.com/Zoxu0928/task-common/logger"
)

// 设置Bean，用于ioc依赖注入
// 支持设置Bean的名称，则在其它Service中可以使用自定义的Bean名称作为变量名称，否则变量名称必须和Bean的默认名称一致
func setBeans() {
	//初始化etcd客户端
	//etcd.NewClient()
}

func main() {
	logger.Debug("test")
}
