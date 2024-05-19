package main

import (
	"fmt"
	"github.com/tidwall/redcon"
)

type MemoryTypeString struct {
	mdb *MemoryDB
}

func (m *MemoryTypeString) memdb() *MemoryDB {
	return m.mdb
}

const (
	MemTypeString = "string"
)

func (m *MemoryTypeString) Type() string {
	return MemTypeString
}

func (m *MemoryTypeString) SizeBytes() uint64 {
	// TODO implement me
	panic("implement me")
}

func (m *MemoryTypeString) Handle(conn redcon.Conn, cmd string, args []string) (mutated bool, err error) {
	switch cmd {
	case "set":
		return
	}
	panic("implement me")
}

func (m *MemoryTypeString) handleSet(conn redcon.Conn, args []string) (bool, error) {
	key := args[0]
	// Check if the string exists
	m.mdb.mu.Lock()
	memString, exists := m.mdb.Keys[key]
	m.mdb.mu.Unlock()

	memType := memString.Type()
	if exists && memType != TypeString {
		conn.WriteError(fmt.Sprintf("key '%s' is a %s, not a  string", key, memType))
		return false, nil
	}

	// TODO implement me
	panic("implement me")
}
