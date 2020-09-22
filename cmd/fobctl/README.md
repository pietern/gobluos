# fobctl

Tool to control [BluOS][bluos] player with [Z-Wave Fibaro KeyFob][keyfob].

[keyfob]: https://manuals.fibaro.com/keyfob/
[bluos]: https://bluos.net/

## Usage

Configure the [Z-Wave to MQTT gateway][zwave2mqtt] to publish all
Z-Wave events to a MQTT broker. Configure `Gateway -> Type` to equal
"ValueID topics" (this means every Z-Wave devices gets its own MQTT
topic prefix). Configure `Gateway -> Payload type` to equal "Entire
Z-Wave value object". Make sure the KeyFob is paired with the Z-Wave
controller.

[zwave2mqtt]: https://github.com/OpenZWave/Zwave2Mqtt

Once done,  edit `fobctl.toml` to reflect your setup, and enjoy!

## License

MIT
