package ws

//Pool keep track of all registerd clients connected
//to the server and provides methods to help manage clients
type Pool struct {
	// Registered clients.
	clients map[string]*Client

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// delegate logic
	delegate WebSocketDelegate
}

func NewPool() *Pool {
	return &Pool{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]*Client),
		delegate:   nil,
	}
}

func (p *Pool) SetDelegate(handler WebSocketDelegate) {
	p.delegate = handler
}

func (p *Pool) GetClient(ID string) *Client {
	return p.clients[ID]
}

func (p *Pool) listen() {
	for {
		select {
		case client := <-p.register:
			p.registerClient(client)
		case client := <-p.unregister:
			p.unregisterClient(client)
		case message := <-p.broadcast:
			p.broadcastMessage(message)
		}
	}
}

func (p *Pool) registerClient(client *Client) {
	p.clients[client.ID] = client
}

func (p *Pool) unregisterClient(client *Client) {
	client.Close()
	if _, ok := p.clients[client.ID]; ok {
		delete(p.clients, client.ID)
	}
}

func (p *Pool) broadcastMessage(message []byte) {
	for _, client := range p.clients {
		client.Send(message)
	}
}
