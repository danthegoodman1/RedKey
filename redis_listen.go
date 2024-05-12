package main

import (
	"github.com/samber/lo"
	"github.com/tidwall/redcon"
	"strings"
)

func redisListen(addr string, memDB *MemoryDB) {
	err := redcon.ListenAndServe(addr,
		func(conn redcon.Conn, args redcon.Command) {
			cmd := strings.ToLower(string(args.Args[0]))
			argStr := lo.Map(args.Args[1:], func(item []byte, index int) string {
				return string(item)
			})
			switch cmd {
			default:
				conn.WriteError("ERR unknown command '" + string(args.Args[0]) + "'")
			case "ping":
				conn.WriteString("PONG")
			case "quit":
				conn.WriteString("OK")
				conn.Close()
			case "set":
				memDB.handleSet(conn, "set", argStr)
			}
		},
		func(conn redcon.Conn) bool {
			// Use this function to accept or deny the connection.
			// log.Printf("accept: %s", conn.RemoteAddr())
			return true
		},
		func(conn redcon.Conn, err error) {
			// This is called when the connection has been closed
			// log.Printf("closed: %s, err: %v", conn.RemoteAddr(), err)
		},
	)
	if err != nil {
		logger.Fatal().Err(err).Msg("redcon.ListenAndServe exited with an error")
	}
}
