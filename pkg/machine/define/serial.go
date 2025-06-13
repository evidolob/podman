package define

import "strings"

type SerialDevice struct {
	//WebSocket host and port in "host:port" form
	WebSocket string `json:",omitempty"`
}

func ParseSerials(config []string) ([]SerialDevice, error) {
	serials := make([]SerialDevice, 0)

	for _, s := range config {
		if strings.HasPrefix(s, "websocket=") {
			hostPort := strings.TrimLeft(s, "websocket=")
			serials = append(serials, SerialDevice{WebSocket: hostPort})
		}
	}

	return serials, nil
}
