package wsapp

const (
	TypeDirectMessage = "message"
)

//Message type a shorthand for map
type Message map[string]string

//MessageCore common fields in every message
type MessageCore struct {
	//Type of the comming message
	Type string `json:"type"`
}

//DirectMessage -
type DirectMessage struct {
	MessageCore `json:",squash"`

	//ID of the message
	MessageID string `json:"message_id,omitempty"`

	//ID of the recipient
	To string `json:"to,omitempty"`

	//ID of the sender
	From string `json:"from,omitempty"`

	//Text content of the message
	Text string `json:"text,omitempty"`

	//Time the message was sent
	Timestamp string `json:"timestamp,omitempty"`
}
