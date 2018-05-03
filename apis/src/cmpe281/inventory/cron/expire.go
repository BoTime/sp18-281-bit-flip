package cron

import (
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx"
	"time"
	"log"
	"github.com/gocql/gocql"
	"cmpe281/inventory/model"
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
	expireTime := time.Now()

	query, names := qb.Select("allocations").
		Columns("store_id", "user_id", "id", "expires").
			Where(qb.Eq("status")).ToCql()
	allocs := make([]model.AllocationDetails, 0)
	q := gocqlx.Query(ctx.Database.Shard1.Query(query), names)
	if err := gocqlx.Select(&allocs, q.Query); err == nil {
		batch := qb.Batch()
		batchData := qb.M{}
		for _, alloc := range allocs {
			prefix := alloc.Id.String()

			batch.AddWithPrefix(prefix, qb.Update("allocations").SetNamed("status", "updated_status").
				Where(qb.Eq("store_id"), qb.Eq("user_id"), qb.Eq("id")).
				If(qb.Eq("status"), qb.Lt("expires")))

			batchData[prefix + ".store_id"] = alloc.StoreId
			batchData[prefix + ".user_id"] = alloc.UserId
			batchData[prefix + ".id"] = alloc.Id
			batchData[prefix + ".expires"] = expireTime
			batchData[prefix + ".updated_status"] = "Expired"
			batchData[prefix + ".status"] = "Unconfirmed"
		}

		updateQuery, updateNames := batch.ToCql()
		q = gocqlx.Query(ctx.Database.Shard1.Query(updateQuery), updateNames).BindMap(batchData)
		if err := q.ExecRelease(); err != nil {
			log.Println("Failed to expire allocations in Shard1")
		}
	}

	allocs = make([]model.AllocationDetails, 0)
	q = gocqlx.Query(ctx.Database.Shard2.Query(query), names)
	if err := gocqlx.Select(&allocs, q.Query); err == nil {
		batch := qb.Batch()
		batchData := qb.M{}
		for _, alloc := range allocs {
			prefix := alloc.Id.String()

			batch.AddWithPrefix(prefix, qb.Update("allocations").SetNamed("status", "updated_status").
				Where(qb.Eq("store_id"), qb.Eq("user_id"), qb.Eq("id")).
				If(qb.Eq("status"), qb.Lt("expires")))

			batchData[prefix + ".store_id"] = alloc.StoreId
			batchData[prefix + ".user_id"] = alloc.UserId
			batchData[prefix + ".id"] = alloc.Id
			batchData[prefix + ".expires"] = expireTime
			batchData[prefix + ".updated_status"] = "Expired"
			batchData[prefix + ".status"] = "Unconfirmed"
		}

		updateQuery, updateNames := batch.ToCql()
		q = gocqlx.Query(ctx.Database.Shard2.Query(updateQuery), updateNames).BindMap(batchData)
		if err := q.ExecRelease(); err != nil {
			log.Println("Failed to expire allocations in Shard2")
		}
	}

	log.Println("Cleaned Finished!")
}