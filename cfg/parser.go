package cfg

import (
	"github.com/Zoxu0928/task-common/db"
	"github.com/Zoxu0928/task-common/http"
	"github.com/Zoxu0928/task-common/logger"
	"github.com/Zoxu0928/task-common/web"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

const (
	DefaultSchedulerLoopInterval time.Duration = 30 * time.Second
	DefaultEtcdFightingPrefix                  = "/distribute_lock/pcd/openapi"
	DefaultFightingInterval      time.Duration = 3 * time.Second
	DefaultEtcdAliveTTL          time.Duration = 1 * time.Second
)

func init() {
	//初始化默认配置
	Configuration = &Config{
		Title:      "Task-Scheduler Default Configuration",
		Env:        EnvConfig{},
		Cluster:    ClusterConfig{},
		Dispatcher: DispatcherConfig{},
		Scheduler:  SchedulerConfig{},
		Log:        logger.Conf{},
		Database:   db.DatabaseConf{},
		Web:        web.WebConf{},
		Etcd:       clientv3.Config{},
		HttpClient: http.HttpClientConf{},
	}
}

//读取配置文件
func (c *Config) Load(configFile *string) error {
	
}
