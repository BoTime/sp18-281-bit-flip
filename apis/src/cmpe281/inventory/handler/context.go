package handler

import (
	"github.com/gocql/gocql"
)

type ShardedDatabaseContext struct {
	Shard1 *gocql.Session
	Shard2 *gocql.Session
}

type RequestContext struct {
	Database ShardedDatabaseContext
}
