package main

import "sync"

type Transaction struct {
	keys      []string
	executeFn func(store map[string]int) bool
}

func NewTransaction(keys []string, executeFn func(map[string]int) bool) *Transaction {
	transaction := &Transaction{
		keys:      keys,
		executeFn: executeFn,
	}

	return transaction
}

func (transaction *Transaction) LinkAndExecuteTransaction(node *Node) bool {

	transactionResult := false

	acquireWG := &sync.WaitGroup{}

	lockedMutexes := make(chan *sync.Mutex, len(transaction.keys))

	tempStore := make(map[string]int)

	for _, val := range transaction.keys {
		acquireWG.Add(1)
		func(key string) { // making this async causes dining philosopher problem
			record := node.store[key]
			record.mutex.Lock()
			lockedMutexes <- record.mutex
			tempStore[key] = record.val
			acquireWG.Done()
		}(val)
	}

	acquireWG.Wait()
	close(lockedMutexes)

	transactionResult = transaction.executeFn(tempStore)

	if transactionResult {
		for key, val := range tempStore {
			node.store[key].val = val
		}
	}

	for mutex := range lockedMutexes {
		mutex.Unlock()
	}

	return transactionResult
}
