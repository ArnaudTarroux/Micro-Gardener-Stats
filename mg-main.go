package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	api "github.com/mg/microgardener/api"
	handler "github.com/mg/microgardener/handler"
	migration "github.com/mg/microgardener/persistence/migration"
)

const (
	worker  = "worker"
	migrate = "migrate"
)

var command string
var version int

func main() {
	flag.StringVar(&command, "command", "worker", "Launch worker or migrations (migrate|void)")
	flag.IntVar(&version, "version", 0, "The version to migrate")
	flag.Parse()

	if command == migrate {
		migration.MigrateDb(version)
		return
	}

	fmt.Println("Starting Micro Gardener stats")

	go api.Init()

	uri := fmt.Sprintf("tcp://%s:%s@mosquitto:1883", os.Getenv("MQTT_USER"), os.Getenv("MQTT_PASSWORD"))

	opts := mqtt.NewClientOptions().AddBroker(uri).SetClientID("microgardener-stats")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(handler.DefaultPublishHandler)
	opts.SetPingTimeout(1 * time.Second)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("mg/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for true {
		select {}
	}
}
