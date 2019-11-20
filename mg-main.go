package main

import (
	"flag"
	"fmt"

	"github.com/mg/microgardener/api"
	"github.com/mg/microgardener/persistence/migration"
	"github.com/mg/microgardener/worker"
)

const (
	workerCommand  = "worker"
	migrateCommand = "migrate"
)

var command string
var version int

func main() {
	flag.StringVar(&command, "command", "worker", "Launch worker or migrations (migrate|void)")
	flag.IntVar(&version, "version", 0, "The version to migrate")
	flag.Parse()

	if command == migrateCommand {
		migration.MigrateDb(version)
		return
	}

	fmt.Println("Starting Micro Gardener stats")

	go api.Init()
	worker.Init()
}
