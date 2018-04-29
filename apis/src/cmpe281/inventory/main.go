package main

import (
	"cmpe281/common/parse"
	"cmpe281/inventory/server"
	"flag"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"time"
	"cmpe281/inventory/cron"
)

func main() {
	var wait time.Duration
	var dbuser, dbpass, dbkeyspace, dbhosts, ip, port string
	var debug bool
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&dbuser, "dbuser", "starbucks",
		"the username for cassandra database connection")
	flag.StringVar(&dbpass, "dbpass", "",
		"the password for cassandra database connection")
	flag.StringVar(&dbkeyspace, "dbkeyspace", "starbucks",
		"the keyspace for cassandra database connection")
	flag.StringVar(&dbhosts, "dbhosts", "54.176.100.87,54.241.192.98",
		"the hosts (comma separated) for cassandra database connection")
	flag.StringVar(&ip, "ip", "0.0.0.0", "ip address to listen on")
	flag.StringVar(&port, "port", "8080", "port to listen on")
	flag.BoolVar(&debug, "debug", false, "run server in debug mode")
	flag.Parse()

	cluster := gocql.NewCluster(parse.SplitCommaSeparated(dbhosts)...)
	cluster.Keyspace = dbkeyspace
	cluster.Timeout = 5 * time.Second
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: dbuser,
		Password: dbpass,
	}

	log.Printf("Connecting to Cassandra...")
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Done")

	apiServer := &server.Server{
		Cassandra: session,
	}

	go func() {
		cleanerContext := cron.CleanerContext{Cassandra: session}
		for range time.NewTicker(time.Minute).C {
			cleanerContext.Cleanup()
		}
	}()

	apiServer.Run(server.Config{
		Ip:   ip,
		Port: port,
		Wait: wait,
	})

	session.Close()
}
