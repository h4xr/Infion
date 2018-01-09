// Copyrights 2018 Saurabh Badhwar. All Rights Reserved.
// The use of this code is governed by MIT License
// which can be found in the LICENSE file.

package server

import (
	"net"
)

const (
	// CLIENTACTIVE Determines the active client
	CLIENTACTIVE = true

	// CLIENTCLOSED Determines the closed client
	CLIENTCLOSED = false
)

// Client structure represents the clients that are connected to our
// platform and are active in nature. The structure helps us maintain
// a common record for the client containing its address and the current state.
type Client struct {
	clientAddr	*net.UDPAddr
	clientState	bool
}

// NewClient creates and returns a new client
func NewClient(address *net.UDPAddr, state bool) *Client {
	return &Client{clientAddr: address, clientState: state}
}

// GetAddr returns the address of the client
func (c Client) GetAddr() *net.UDPAddr {
	return c.clientAddr
}

// GetState returns the state of the client
func (c Client) GetState() bool {
	return c.clientState
}

//SetState sets the current state of the client
func (c *Client) SetState(state bool) {
	c.clientState = state
}

// ClientPool structure holds the clients and the associated topics the
// the client has subscribed to.
type ClientPool struct {
	pool	map[string][]Client
	clientList	map[Client]bool
}

// NewClientPool returns a new client pool in which new clients can be added
func NewClientPool() *ClientPool {
	return &ClientPool{pool: make(map[string][]Client), clientList: make(map[Client]bool)}
}

// AddClient adds a new client to the client pool
func (cp *ClientPool) AddClient(topic string, client Client) {
	if _, ok := cp.pool[topic]; ok {
		cp.pool[topic] = append(cp.pool[topic], client)
	} else {
		cp.pool[topic] = []Client{client}
	}
	cp.clientList[client] = true
}

// GetClients returns the clients subscribed to the provided topic
func (cp ClientPool) GetClients(topic string) []Client {
	if clients, ok := cp.pool[topic]; ok {
		return clients
	}
	return nil
}