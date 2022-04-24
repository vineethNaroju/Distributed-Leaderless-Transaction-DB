package main

import (
	"fmt"
	"sync"
	"time"
)

type initRecord struct {
	key string
	val int
}

func Demo() {
	db := NewNode("earth")

	records := []initRecord{
		{"rose", 100},
		{"john", 20},
	}

	for _, r := range records {
		go db.CreateRecord(r.key, r.val)
	}

	time.Sleep(1 * time.Second)

	transferAmount := 10

	transactionKeys := []string{"rose", "john"}

	executeFn := func(store map[string]int) bool {
		rose, john := "rose", "john"

		roseAmt := store[rose]
		johnAmt := store[john]

		fmt.Println(roseAmt, johnAmt)

		if roseAmt-transferAmount >= 0 {
			johnAmt += transferAmount
			roseAmt -= transferAmount

			store[rose] = roseAmt
			store[john] = johnAmt

			return true
		}

		return false
	}

	simpleTransaction := NewTransaction(transactionKeys, executeFn)

	wg := &sync.WaitGroup{}

	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("transaction status", db.ExecuteTransaction(simpleTransaction))
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(db.Get("rose"), db.Get("john"))
}
