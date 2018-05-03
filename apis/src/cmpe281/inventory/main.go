package main

import (
	"cmpe281/common/parse"
	"cmpe281/inventory/server"
	"flag"
	"github.com/gocql/gocql"
	"log"
	"time"
	"cmpe281/inventory/cron"
)

func main() {
	var wait time.Duration
	var dbuser, dbpass, dbkeyspace, dbshard1, dbshard2, ip, port string
	var debug bool
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&dbuser, "dbuser", "starbucks",
		"the username for cassandra database connection")
	flag.StringVar(&dbpass, "dbpass", "",
		"the password for cassandra database connection")
	flag.StringVar(&dbkeyspace, "dbkeyspace", "starbucks",
		"the keyspace for cassandra database connection")
	flag.StringVar(&dbshard1, "dbshard1", "54.176.100.87,54.241.192.98",
		"the hosts (comma separated) for first cassandra database shard")
	flag.StringVar(&dbshard2, "dbshard2", "13.57.76.150",
		"the hosts (comma separated) for second cassandra database shard")
	flag.StringVar(&ip, "ip", "0.0.0.0", "ip address to listen on")
	flag.StringVar(&port, "port", "8080", "port to listen on")
	flag.BoolVar(&debug, "debug", false, "run server in debug mode")
	flag.Parse()

	log.Printf("Connecting to Cassandra... Shard1 | ")
	shard1config := generateBaseCassandraConfig(dbshard1, dbuser, dbpass, dbkeyspace)
	shard1, err := shard1config.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Done ... Shard2 | ")

	shard2config := generateBaseCassandraConfig(dbshard2, dbuser, dbpass, dbkeyspace + "2")
	shard2, err := shard2config.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Done")

	apiServer := &server.Server{
		Database: server.ShardedDatabaseContext{
			Shard1: shard1,
			Shard2: shard2,
		},
	}

	go func() {
		cleanerContext := cron.CleanerContext{Database: cron.ShardedDatabaseContext{
			Shard1: shard1,
			Shard2: shard2,
		}}
		for range time.NewTicker(time.Minute).C {
			cleanerContext.Cleanup()
		}
	}()

	apiServer.Run(server.Config{
		Ip:   ip,
		Port: port,
		Wait: wait,
	})

	apiServer.Database.Shard1.Close()
	apiServer.Database.Shard2.Close()
}

func generateBaseCassandraConfig(dbhosts string, dbuser string , dbpass string, dbkeyspace string) (*gocql.ClusterConfig) {
	cluster := gocql.NewCluster(parse.SplitCommaSeparated(dbhosts)...)
	cluster.Keyspace = dbkeyspace
	cluster.Timeout = 5 * time.Second
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: dbuser,
		Password: dbpass,
	}
	cluster.ReconnectionPolicy = &gocql.ConstantReconnectionPolicy{
		MaxRetries: 3,
		Interval: 1 * time.Second,
	}
	cluster.RetryPolicy = &gocql.DowngradingConsistencyRetryPolicy{
		ConsistencyLevelsToTry: []gocql.Consistency {
			gocql.Quorum,
			gocql.LocalQuorum,
			gocql.One,
		},
	}
	cluster.IgnorePeerAddr = true
	cluster.DisableInitialHostLookup = true
	return cluster
}