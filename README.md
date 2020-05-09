# samsung-tv-ex-link-mqtt-client
Control the TV using MQTT over the TV's EX-Link port

* mqtt.go - An mqtt client that listens for commands like "on" and does actions like turn on the tv when received.
* on.go - Immediately turn on the tv
* hdmi.go - Immediately switches to input HDMI1

Connect a computer/raspberry pi to the TV through a serial to EX-Link adapter (through a USB to serial adapter), compile and run.

Edit and add `tv.service` to systemd to start on boot.
