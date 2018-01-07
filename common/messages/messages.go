// Copyrights 2018 Saurabh Badhwar. All Rights Reserved.
// The use of this code is governed by MIT License
// which can be found in the LICENSE file.

// Package messages provide functionality related to the handling of the
// the incoming messages and how to deal with them based on their type.
package messages

import (
	"encoding/json"
)

// MessageType defines the different set of messages which are supported by us
type MessageType uint8

// MessageTypes supported by the messages engine
var MessageTypes = map[MessageType]string{
	// TEST message type. Used for testing the connection
	0x00: "TEST",
	// PING message type. Used to check liveliness of the recepient
	0x01: "PING",
	// SHUTDOWN message type. Used for shutting down the client.
	0x02: "SHUTDOWN",
	// CLOSE message type. Used to disconnect the client.
	0x03: "CLOSE",
	// MAINTAINANCE message type. Used to send client into maintainance.
	0x04: "MAINTAINANCE",
	// CONTROL message type. Used to make the client switch to operational mode
	// from maintainance mode.
	0x05: "CONTROL",
	// REGISTER message type. Used to register client to broker.
	0x06: "REGISTER",
}

// Message defines a structure for incoming and outgoing messages
// to be used by the UDP server. 
type Message struct {
	// MsgType defines the type of message being sent
	// Infion already has a set of predefined set of messages
	// that is uses internally to manage the nodes.
	// The message type is represented by a 8 bit unsigned int
	// which is used for efficiently representing a lot of
	// different messages that can be exchanged between the nodes.
	MsgType		MessageType	`json:"type"`

	// Checksum is a md5 checksum of the payload that is being sent
	// to the recieveing party. The recieving party can use this
	// checksum to verify if the content that is sent is valid
	// or not.
	Checksum		string	`json:"checksum"`

	// Payload defines the data that needs to be sent to the receiver
	// which has to process it. How the payload is processed is
	// completely dependent upon the type handler being used by the
	// recieving end.
	Payload			string	`json:"payload"`
}

// NewMessage defines and returns a new message which can be processed by the
// broker or clients.
func NewMessage(mtype MessageType, payload string) *Message {
	message := new(Message)
	message.MsgType = mtype
	message.Payload = payload
	message.Checksum = generateChecksum(message.Payload)

	return message
}

// FromJSON parses the JSON and returns a Message type
func FromJSON(data []byte) (*Message, error) {
	message := new(Message)
	err := json.Unmarshal(data, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

// ToJSON returns the JSON formatted representation of the structure
func (m Message) ToJSON() ([]byte, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetType returns the type of the message
func (m Message) GetType() MessageType {
	return m.MsgType
}

// GetPayload returns the payload from the message
func (m Message) GetPayload() string {
	return m.Payload
}

// ValidateType returns a bool indicating if the message type provided is
// valid or not.
func (m Message) ValidateType() bool {
	if _, ok := MessageTypes[m.MsgType]; ok {
		return true
	}
	return false
}

// VerifyMessage verifies if the message recieved is valid or not
func (m Message) VerifyMessage() bool {
	checksum := generateChecksum(m.Payload)
	if checksum == m.Checksum {
		return true
	}
	return false
}

