package flow

import (
	"sync"
	"sync/atomic"
)

type WorkFlow struct {
	done     chan struct{}
	doneOnce *sync.Once
	root     *Node
	End      *Node
	edges    []*Edge
}

type Edge struct {
	FromNode *Node
	ToNode   *Node
}

type Node struct {
	Dependencies []*Edge
	DepCompleted int32
	Task         Runnable
	Children     []*Edge
}

func (n *Node) Execute(i interface{}) {
	if !n.dependencyHasDone() {
		return
	}

	if n.Task != nil {
		n.Task.Run(i)
	}

	for _, child := range n.Children {
		child.ToNode.Execute(i)
	}
}

func (n *Node) dependencyHasDone() bool {
	if n.Dependencies == nil {
		return true
	}

	atomic.AddInt32(&n.DepCompleted, 1)

	return int32(len(n.Dependencies)) == n.DepCompleted
}

type Runnable interface {
	Run(interface{})
}

func NewNode(Task Runnable) *Node {
	return &Node{
		Task: Task,
	}
}

func AddEdge(from *Node, to *Node) *Edge {
	edge := &Edge{
		FromNode: from,
		ToNode:   to,
	}

	//该条边是from节点的出边
	from.Children = append(from.Children, edge)
	//该条边是to节点的入边
	to.Dependencies = append(to.Dependencies, edge)

	return edge
}
