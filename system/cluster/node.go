package cluster

import "fmt"

type INode interface {
	GetName() string
	GetIP() string
	GetNameSpace() string
	GetDetail() string
}

type Node struct {
	nodeName  string // 当前节点名称
	name      string // 当前pod名称
	ip        string // 当前节点(pod)的ip地址
	nameSpace string // 当前namespace
}

func (n *Node) GetName() string {
	return n.name
}

func (n *Node) GetIP() string {
	return n.ip
}

func (n *Node) GetNameSpace() string {
	return n.nameSpace
}

func (n *Node) GetDetail() string {
	return fmt.Sprintf("{node=%s name=%s ip=%s namespace=%s}", n.nodeName, n.name, n.ip, n.nameSpace)
}
