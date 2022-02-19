package main

import (
	"github.com/Zoxu0928/task-common/basic"
	ctxBean "github.com/Zoxu0928/task-common/bean"
	"github.com/Zoxu0928/task-common/distribute_mutex"
	"github.com/Zoxu0928/task-common/etcd"
	"github.com/Zoxu0928/task-common/logger"
	"github.com/Zoxu0928/task-scheduler/cfg"
)

func init() {

}

// 设置Bean，用于ioc依赖注入
// 支持设置Bean的名称，则在其它Service中可以使用自定义的Bean名称作为变量名称，否则变量名称必须和Bean的默认名称一致
func setBeans() {
	//初始化etcd客户端
	//设置公共资源Bean
	etcdClient, err := etcd.NewClient(cfg.Configuration.Etcd)
	//无法容忍错误
	basic.Must(nil, err)
	ctxBean.AddBean(etcdClient)
	//初始化ETCD分布式锁
	//ETCD目录
	prefix := cfg.Configuration.Cluster.EtcdFightingPrefix
	//锁的keepalive周期
	ttl := cfg.Configuration.Cluster.EtcdKeepAliveTTL
	//创建分布式锁
	distribute_mutex.DefaultDistributeMutex, err = distribute_mutex.NewETCDDistributeMutex(etcdClient, prefix, ttl.Duration)
	basic.Must(nil, err)
	ctxBean.AddBean(distribute_mutex.DefaultDistributeMutex)
	//初始化集群
}

func main() {
	setBeans()
	logger.Debug("test")
}
