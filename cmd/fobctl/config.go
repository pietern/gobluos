package main

type MQTT struct {
	// Broker is a fully qualified URL to the MQTT broker.
	// For example: "tcp://192.168.1.3:1883".
	Broker string `toml:"broker"`

	// ClientID is the MQTT client ID. Default: "fobctl".
	ClientID string `toml:"client_id"`

	// KeyFobTopic is the MQTT topic where the Fibaro KeyFob
	// messages are published. For example: "zwave/19/91/1/#"
	KeyFobTopic string `toml:"key_fob_topic"`
}

type Configuration struct {
	// PlayerURL is the fully qualified URL to the BluOS player.
	// For example: "http://192.168.1.2:11000".
	PlayerURL string `toml:"player_url"`

	// MQTT contains all MQTT related configuration.
	MQTT MQTT
}
