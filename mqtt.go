package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	//"strconv"
	"syscall"

	MQTT "github.com/eclipse/paho.mqtt.golang"
        "github.com/mgoff/go-samsung-exlink"
        "log"
)

func sendCommand(command string) {

	// open the connection to the EX-Link device
	device, err := exlink.Open("/dev/ttyUSB0")
	if err != nil {
		log.Fatal(err)
	}

	// close the connection at the end
	defer device.Close()

	switch command {
		case "on":
			err = device.PowerOn()
		case "off":
			err = device.PowerOff()
		case "hdmi1":
			err = device.SourceHDMI1()
		case "hdmi2":
			err = device.SourceHDMI2()
//		case "hdmi3":
//			err = device.SourceHDMI3()
//		case "hdmi4":
//			err = device.SourceHDMI4()
	}

	if err != nil {
		log.Fatal(err)
	}
}

func onMessageReceived(client MQTT.Client, message MQTT.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
	sendCommand(string(message.Payload()[:])) 
}

func main() {
	//MQTT.DEBUG = log.New(os.Stdout, "", 0)
	//MQTT.ERROR = log.New(os.Stdout, "", 0)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	//hostname, _ := os.Hostname()

	server := flag.String("server", "tcp://192.168.0.2:1883", "The full url of the MQTT server to connect to ex: tcp://127.0.0.1:1883")
	topic := flag.String("topic", "devices/tv/command", "Topic to subscribe to")
	qos := flag.Int("qos", 0, "The QoS to subscribe to messages at")
	clientid := flag.String("clientid", "tv", "A clientid for the connection")
	username := flag.String("username", "", "A username to authenticate to the MQTT server")
	password := flag.String("password", "", "Password to match username")
	flag.Parse()

	connOpts := MQTT.NewClientOptions().AddBroker(*server).SetClientID(*clientid).SetCleanSession(true)
	if *username != "" {
		connOpts.SetUsername(*username)
		if *password != "" {
			connOpts.SetPassword(*password)
		}
	}

	connOpts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(*topic, byte(*qos), onMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}

	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", *server)
	}

	<-c
}
