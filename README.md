# RedKey
Redis compatible in-memory database that is persisted, and replicated via RedPanda (like AWS MemoryDB but open source)

## Persistence

Like AWS MemoryDB, RedKey persists all mutations durably before responding to a command.

RedKey uses 2 levels of persistence:

The first is the log, powered by RedPanda. All mutations are written to the log, and the command does not respond until RedPanda acknowledges durable operation. This is the equivalent to Redis AOF.

The second is a local Badger database (key value). This is the equivalent to Redis RDB snapshots.

On an interval, the KV is snapshotted and written to S3 for faster recovery and replica creation. RedKey manages a transaction which updates the mutations in the KV async from the response, also noting the latest offset written to the log. These operations must be done sequentially, but they are buffered and done in batches to ensure high throughput. These mutations do not need to succeed, and on transaction failure will crash the node (after writing to the log). This ensures that we do not slow down command processing, while also ensuring we have fast recovery times with a snapshot.

## Recovery

First takes snapshot, writes in mem, then consumes from log.