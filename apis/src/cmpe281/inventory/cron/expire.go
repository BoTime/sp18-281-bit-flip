package cron

import (
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx"
	"time"
	"log"
	"github.com/gocql/gocql"
)

type ShardedDatabaseContext struct {
	Shard1 *gocql.Session
	Shard2 *gocql.Session
}

type CleanerContext struct {
	Database ShardedDatabaseContext
}

func (ctx *CleanerContext) Cleanup() {
	// Don't put inventory back for now, just expire the entries.
	// TODO(bbamsch): Restore inventory

	log.Println("Cleanup Started...")

	query, names := qb.Update("allocations").SetNamed("status", "updated_status").Where(qb.Eq("status"), qb.Lt("expires")).ToCql()
	q := gocqlx.Query(ctx.Database.Shard1.Query(query), names).BindMap(qb.M{
		"updated_status": "Expired",
		"status": []string{"Unconfirmed"},
		"expires": time.Now(),
	})

	if err := q.ExecRelease(); err != nil {
		log.Println("Failed to expire allocations from Shard1")
		return
	}

	q = gocqlx.Query(ctx.Database.Shard2.Query(query), names).BindMap(qb.M{
		"updated_status": "Expired",
		"status": []string{"Unconfirmed"},
		"expires": time.Now(),
	})

	if err := q.ExecRelease(); err != nil {
		log.Println("Failed to expire allocations from Shard2")
		return
	}

	log.Println("Cleaned Finished!")
}
