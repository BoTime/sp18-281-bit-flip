package cron

import (
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx"
	"github.com/gocql/gocql"
	"time"
	"log"
)

type CleanerContext struct {
	Cassandra *gocql.Session
}

func (ctx *CleanerContext) Cleanup() {
	// Don't put inventory back for now, just expire the entries.
	// TODO(bbamsch): Restore inventory

	log.Println("Cleanup Started...")

	query, names := qb.Update("allocations").SetNamed("status", "updated_status").Where(qb.Eq("status"), qb.Lt("expires")).ToCql()
	q := gocqlx.Query(ctx.Cassandra.Query(query), names).BindMap(qb.M{
		"updated_status": "Expired",
		"status": []string{"Unconfirmed"},
		"expires": time.Now(),
	})

	if err := q.ExecRelease(); err != nil {
		log.Println("Failed to expire allocations")
		return
	}

	log.Println("Cleaned Finished!")
}
