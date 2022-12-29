package main

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/karim-w/stdlib"
)

type DatabaseConfig struct {
	SQL []DBOptions       `json:"sql",yaml:"sql"`
	CQL []CassandraConfig `json:"cql",yaml:"cql"`
}

type DBOptions struct {
	DatabaseDSN   string   `json:"databaseDsn" yaml:"databaseDsn"`
	DatabaseType  string   `json:"databaseType" yaml:"databaseType"`
	DatabaseNames []string `json:"databaseNames" yaml:"databaseNames"`
}

type CassandraConfig struct {
	URLs              []string `json:"urls" yaml:"urls"`
	UserName          string   `json:"username" yaml:"username"`
	Password          string   `json:"password" yaml:"password"`
	Class             string   `json:"class" yaml:"class"`
	ReplicationFactor int      `json:"replicationFactor" yaml:"replicationFactor"`
	KeySpaces         []string `json:"keyspaces" yaml:"keyspaces"`
}

const (
	CASSANDRA_KEYSPACE_INIT = `CREATE KEYSPACE IF NOT EXISTS %s WITH REPLICATION = {
  	'class' : '%s',
  	'replication_factor' : '%d'
  };`
)

func (db *DatabaseConfig) Setup() {
	// setup the databases
	for _, db := range db.SQL {
		dbctx := stdlib.NativeDatabaseProvider(
			db.DatabaseType,
			db.DatabaseDSN,
		)
		for _, dbname := range db.DatabaseNames {
			_, err := dbctx.Exec("CREATE DATABASE " + dbname)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	for _, db := range db.CQL {
		auth := gocql.PasswordAuthenticator{
			Username: db.UserName,
			Password: db.Password,
		}

		// - creating main session
		cls := gocql.NewCluster(db.URLs...)
		cls.Authenticator = auth
		sess, err := cls.CreateSession()
		if err != nil {
			panic(err)
		}
		for _, keyspace := range db.KeySpaces {
			qry := fmt.Sprintf(CASSANDRA_KEYSPACE_INIT, keyspace, db.Class, db.ReplicationFactor)
			err := sess.Query(qry).Exec()
			if err != nil {
				fmt.Println(err)
			}

		}
	}
}
