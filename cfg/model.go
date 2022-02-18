package cfg

import (
	"github.com/Zoxu0928/task-common/basic"
	"github.com/Zoxu0928/task-common/db"
	"github.com/Zoxu0928/task-common/http"
	"github.com/Zoxu0928/task-common/logger"
	"github.com/Zoxu0928/task-common/web"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var Configuration *Config

// EnvConfig 环境配置
type EnvConfig struct {
	// 禁用优雅停止
	DisableGracefullyStopped bool `toml:"disable_gracefully_stopped"`
}

//集群配置
type ClusterConfig struct {
	// 集群选取的周期
	// 建议设置相对较小的值
	// 默认3s
	FightingInterval basic.Duration `toml:"fighting_interval" validate:"gte=1,lte=10"`
	// 集群通过ETCD进行选举，需要定义分布式锁的prefix
	EtcdFightingPrefix string         `toml:"etcd_fighting_prefix" validate:"required,startswith=/distribute_lock/"`
	EtcdKeepAliveTTL   basic.Duration `toml:"etcd_keep_alive_ttl" validate:"gte=1,lte=10"`
}

// SchedulerConfig 调度器配置
type SchedulerConfig struct {
	// 调度周期
	// 默认30s
	LoopInterval basic.Duration `toml:"loop_interval" validate:"gte=3,lte=300"`
}

// DispatcherConfig 任务分发器配置
type DispatcherConfig struct {
}

type Config struct {
	Title      string `validate:"required"`
	Env        EnvConfig
	Cluster    ClusterConfig
	Dispatcher DispatcherConfig
	Scheduler  SchedulerConfig
	Log        logger.Conf
	Database   db.DatabaseConf
	Web        web.WebConf
	Etcd       clientv3.Config
	HttpClient http.HttpClientConf
}
