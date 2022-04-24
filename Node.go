package main

import "sync"

type Node struct {
	store         map[string]*Record
	creationMutex *sync.Mutex
}

func NewNode(name string) *Node {
	node := &Node{
		store:         make(map[string]*Record),
		creationMutex: &sync.Mutex{},
	}

	return node
}

func (node *Node) CreateRecord(key string, val int) bool {

	node.creationMutex.Lock()

	defer node.creationMutex.Unlock()

	if _, ok := node.store[key]; ok {
		return false
	}

	node.store[key] = NewRecord(val)

	return false
}

func (node *Node) ExecuteTransaction(transaction *Transaction) bool {
	return transaction.LinkAndExecuteTransaction(node)
}

func (node *Node) Get(key string) int {
	record := node.store[key]

	if record == nil {
		return -1
	}

	record.mutex.Lock()
	defer record.mutex.Unlock()

	return record.val
}
