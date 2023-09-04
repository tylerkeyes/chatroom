package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gocql/gocql"
)

// path to config file
var configPath = "./config/config.env"

// configuration variables
var cassandraClusterIp = "127.0.0.1:9042"

// globals
var router *chi.Mux

type Server struct {
	router *chi.Mux
	db     *gocql.Session
}

// TODO: create global variable for db access

func main() {
	loadConfig()

	server := initServer()
	fmt.Println("Created server")
	defer server.db.Close()

	var err error

	err = http.ListenAndServe(":3333", server.router)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func loadConfig() {
	var configs = make(map[string]string)
	
	file, err := os.Open(configPath)
	if err != nil {
		log.Printf("error opening config file: %s\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), "=")
		configs[text[0]] = text[1]
	}

	if err := scanner.Err(); err != nil {
		log.Printf("error scanning config file: %s\n", err)
	}

	// TODO: read config options into variables
	fmt.Printf("cluster ip: %v\n", cassandraClusterIp)
	cassandraClusterIp = configs["cassandra_ip"]
	fmt.Printf("cluster ip: %v\n", cassandraClusterIp)
}

func initServer() *Server {
	fmt.Println("Initialize server")

	// init router
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)

	// init messages db
	db := initDbConnection()

	s := Server{
		router: router,
		db:     db,
	}

	// init server routes
	s.Routes()

	return &s
}

func initDbConnection() *gocql.Session {
	// need to find a way to know the IP address of the machine running cassandra
	cluster := gocql.NewCluster(cassandraClusterIp)
	fmt.Println("Creating cluster")
	cluster.Keyspace = "chatbot"
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	fmt.Printf("Creating cluster with Keyspace: %s, Consistency: %s, ProtoVersion: %d\n", cluster.Keyspace, cluster.Consistency, cluster.ProtoVersion)

	// init Cassandra DB session (note: close not defered, need to close session automatically)
	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Printf("Error when connecting to Cassandra DB cluster: %v", err)
		// return nil when database connection not created, better to start server and repair db connection later
		return nil
	}

	return session
}
