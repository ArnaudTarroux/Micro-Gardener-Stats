package main

import (
	"fmt"
	"os"
	"time"

	handler "github.com/mg/microgardener/handler"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	fmt.Println("Starting Micro Gardener stats")

	opts := mqtt.NewClientOptions().AddBroker("tcp://mosquitto:1883").SetClientID("microgardener-stats")
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
