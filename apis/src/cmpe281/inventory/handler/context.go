package handler

import "github.com/gocql/gocql"

type RequestContext struct {
	Cassandra *gocql.Session
}
