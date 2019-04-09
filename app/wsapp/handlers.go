package wsapp

import (
	"encoding/json"
	"fmt"

	"github.com/TarekkMA/GeoChat/pkg/ws"
)

func (d *Delegate) handleDirectMessage(msg DirectMessage, client *ws.Client) error {

	to := d.pool.GetClient(msg.To)
	if to == nil {
		return fmt.Errorf("client with id %s was not found", msg.To)
	}

	msg.MessageID = NewMessageID()
	msg.Timestamp = TimeNowNanoStr()
	msg.From = client.ID

	b, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	to.Send(b)
	return nil
}
