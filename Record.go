package main

import "sync"

type Record struct {
	mutex *sync.Mutex
	val   int
}

func NewRecord(val int) *Record {
	record := &Record{
		mutex: &sync.Mutex{},
		val:   val,
	}

	return record
}
