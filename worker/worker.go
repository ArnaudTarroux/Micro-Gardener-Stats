package worker

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mg/microgardener/worker/handler"
)

func Init() {
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
