package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pietern/gobluos"
)

type Message struct {
	ValueID    string   `json:"value_id"`
	NodeID     int      `json:"node_id"`
	ClassID    int      `json:"class_id"`
	Type       string   `json:"type"`
	Genre      string   `json:"genre"`
	Instance   int      `json:"instance"`
	Index      int      `json:"index"`
	Label      string   `json:"label"`
	Units      string   `json:"units"`
	Help       string   `json:"help"`
	ReadOnly   bool     `json:"read_only"`
	WriteOnly  bool     `json:"write_only"`
	Min        int      `json:"min"`
	Max        int      `json:"max"`
	IsPolled   bool     `json:"is_polled"`
	Values     []string `json:"values"`
	Value      string   `json:"value"`
	LastUpdate int64    `json:"lastUpdate"`
}

type Handler struct {
	client *gobluos.Client
	logger *log.Logger
}

func NewHandler(address string, logger *log.Logger) *Handler {
	client := gobluos.NewClient(address)
	return &Handler{client, logger}
}

func (h *Handler) HandleMessage(client mqtt.Client, raw mqtt.Message) {
	var msg Message
	if err := json.Unmarshal(raw.Payload(), &msg); err != nil {
		panic(err)
	}

	switch msg.Index {
	case 1:
		h.HandleStop(msg)
	case 2:
		h.HandleStart(msg)
	case 3:
		h.HandlePrevious(msg)
	case 4:
		h.HandleNext(msg)
	case 5:
		h.HandleVolumeDown(msg)
	case 6:
		h.HandleVolumeUp(msg)
	default:
		panic("Unknown button index")
	}
}

func (h *Handler) HandleStop(msg Message) {
	if msg.Value != "Pressed 1 Time" {
		return
	}

	_, err := h.client.Pause()
	if err != nil {
		panic(err)
	}

	h.logger.Printf("Pressed stop")
}

func (h *Handler) HandleStart(msg Message) {
	if msg.Value != "Pressed 1 Time" {
		return
	}

	_, err := h.client.Play()
	if err != nil {
		panic(err)
	}

	h.logger.Printf("Pressed start")
}

func (h *Handler) HandlePrevious(msg Message) {
	if msg.Value != "Pressed 1 Time" {
		return
	}

	_, err := h.client.Back()
	if err != nil {
		panic(err)
	}

	h.logger.Printf("Pressed back")
}

func (h *Handler) HandleNext(msg Message) {
	if msg.Value != "Pressed 1 Time" {
		return
	}

	_, err := h.client.Skip()
	if err != nil {
		panic(err)
	}

	h.logger.Printf("Pressed skip")
}

func (h *Handler) HandleVolumeUp(msg Message) {
	if msg.Value != "Pressed 1 Time" {
		return
	}

	h.AdjustVolume(+5)
}

func (h *Handler) HandleVolumeDown(msg Message) {
	if msg.Value != "Pressed 1 Time" {
		return
	}

	h.AdjustVolume(-5)
}

func (h *Handler) AdjustVolume(diff int) {
	v, err := h.client.Volume()
	if err != nil {
		panic(err)
	}

	level := v.Level + diff
	if level < 0 {
		level = 0
	}
	if level > 100 {
		level = 100
	}

	h.logger.Printf("Setting volume to %d", level)
	_, err = h.client.SetVolume(level)
	if err != nil {
		panic(err)
	}
}

func main() {
	var path string
	flag.StringVar(&path, "config", "", "Path to configuration file")
	flag.Parse()

	var config Configuration
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	handler := NewHandler(config.PlayerURL, logger)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.MQTT.Broker)
	opts.SetClientID(config.MQTT.ClientID)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetReconnectingHandler(func(client mqtt.Client, opts *mqtt.ClientOptions) {
		logger.Fatal("No support for MQTT reconnects...")
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	log.Printf("Subscribing to MQTT topic %#v", config.MQTT.KeyFobTopic)
	token := client.Subscribe(config.MQTT.KeyFobTopic, 1, handler.HandleMessage)
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	log.Printf("Ready...")

	// Halt
	select {}
}
