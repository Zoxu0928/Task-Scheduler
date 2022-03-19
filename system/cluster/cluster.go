package cluster

import (
	"context"
	"errors"
	"github.com/Zoxu0928/task-common/distribute_mutex"
	"sync"
	"time"
)

type Role string

const (
	RoleMaster         Role   = "master"
	RoleSlave          Role   = "slave"
	DefaultClusterName string = "default"
)

var (
	MutexError     = errors.New("[cluster] distribution mutex builder is nil")
	DefaultCluster *Cluster
)

// Cluster 基于ETCD分布式锁实现的集群模式
type Cluster struct {
	mutex sync.Mutex
	name  string
	// 集群角色，启动时根据分布式锁的获取情况赋值
	role Role

	// 分布式锁载体
	distributeMutex  distribute_mutex.IMutex
	fightingInterval time.Duration

	// 协程间通信，控制协程的结束
	ctx context.Context
	// 通知ctx结束
	cancel context.CancelFunc

	started bool
	// 2/2 集群是否已经退出
	stopped bool
	// 1/2 集群是否已经选举完成
	readied bool

	// 节点信息
	node *Node // 当前节点
}
