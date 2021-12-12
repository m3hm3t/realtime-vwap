package subscriber

import ws "golang.org/x/net/websocket"

type Subscriber interface {
	Subscribe(conn *ws.Conn, pairs []string) error
}
