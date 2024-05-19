package main

import (
	"sync"
)

type (
	MemoryDB struct {
		Keys map[string]IMemoryDataType
		mu   *sync.RWMutex
	}
)
