package main

import (
	"fmt"
	"github.com/tidwall/redcon"
	"time"
)

func (mdb *MemoryDB) parseSet(conn redcon.Conn, args []string) {
	// TODO implement me
	mdb.handleSet(conn, args[0], args[1], "", false, 0)
}

func (mdb *MemoryDB) handleSet(conn redcon.Conn, key, value, nxORxx string, returnVal bool, expiry time.Duration) {
	// Check if the string exists
	mdb.mu.Lock()
	memString, exists := mdb.Keys[key]
	mdb.mu.Unlock()

	memType := memString.Type()
	if exists && memType != TypeString {
		conn.WriteError(fmt.Sprintf("key '%s' is a %s, not a  string", key, memType))
		return
	}

	// TODO implement me
}

type MemoryTypeString struct {
}

func (m MemoryTypeString) Handle(conn redcon.Conn, cmd redcon.Command) (mutated bool, err error) {
	// TODO implement me
	panic("implement me")
}

func (m MemoryTypeString) Type() string {
	// TODO implement me
	panic("implement me")
}

func (m MemoryTypeString) SizeBytes() uint64 {
	// TODO implement me
	panic("implement me")
}
