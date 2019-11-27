package publisher

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mg/microgardener/worker/handler"
)

func PublishTo(topic string, payload interface{}) {

	uri := fmt.Sprintf("tcp://%s:%s@mosquitto:1883", os.Getenv("MQTT_USER"), os.Getenv("MQTT_PASSWORD"))
	opts := mqtt.NewClientOptions().AddBroker(uri).SetClientID("microgardener-publisher")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(handler.DefaultPublishHandler)
	opts.SetPingTimeout(1 * time.Second)
	client := mqtt.NewClient(opts)

	defer client.Disconnect(0)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	if token := client.Publish(topic, 0, false, payload); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
