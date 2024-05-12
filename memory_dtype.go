package main

import "github.com/tidwall/redcon"

type (
	IMemoryDataType interface {
		// `mutated` is whether this command modified data.
		// If the command was a no-op (e.g. delete something that didn't exist) it is safe
		// to return `false`, however it is also safe to return `true` no matter what if
		// if it not known. This determins whether to persist it to the RedPanda log.
		//
		// Each `Handle` is responsible for updating the local KV snapshot (async)
		Handle(conn redcon.Conn, cmd redcon.Command) (mutated bool, err error)

		// Type returns the data structure type, checked on initial command recv
		Type() string

		SizeBytes() uint64
	}
)

const (
	TypeString = "string"
)
