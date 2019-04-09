package wsapp

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mitchellh/mapstructure"

	"github.com/TarekkMA/GeoChat/pkg/ws"
)

func NewDelegate(pool *ws.Pool) *Delegate {
	return &Delegate{
		pool: pool,
	}
}

type Delegate struct {
	pool *ws.Pool
}

func (d *Delegate) HandleMessage(client *ws.Client, rawMessage []byte) {
	message, err := parseMessage(rawMessage)
	if err != nil {
		log.Printf("[ERR] [MessageHandler]: message parsing %s %+v", err, err)
		return
	}
	if err := d.routeToHandler(message, client); err != nil {
		log.Printf("[ERR] [MessageHandler]: routeToHandler %s %+v", err, err)
		return
	}
}

func (d *Delegate) HandleRegister(c *ws.Client, r *http.Request) bool {
	c.Send([]byte("ID:[" + c.ID + "]"))
	return true
}

func parseMessage(rawMessage []byte) (Message, error) {
	var parsed Message
	err := json.Unmarshal(rawMessage, &parsed)
	return parsed, err
}

func (d *Delegate) routeToHandler(message Message, client *ws.Client) error {

	switch message["type"] {
	case TypeDirectMessage:
		var m DirectMessage
		if err := decodeMap(message, &m); err != nil {
			return err
		}
		return d.handleDirectMessage(m, client)
	}
	return nil
}

func decodeMap(input Message, result interface{}) error {
	d, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		//reusing the same tag as the json tag
		//to reduce duplication
		TagName: "json",
		Result:  result,
	})
	if err != nil {
		return err
	}
	return d.Decode(input)
}
