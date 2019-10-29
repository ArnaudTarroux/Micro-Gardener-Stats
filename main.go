package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	handler "./handler"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	fmt.Println("Starting Micro Gardener stats")

	http.Handle("/metrics", promhttp.Handler())

	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("microgardener-stats")
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

	}
}
