// Copyrights 2018 Saurabh Badhwar. All Rights Reserved.
// The use of this code is governed by MIT License
// which can be found in the LICENSE file.

package messages

import (
	"fmt"
)

// MessageHandlers defines a type to register different message
// handlers respective of the types they are supposed to handle.
// For a new message handler, they have to provide the type of message
// they can handle and the handler function that should be called
// with the payload.
type MessageHandlers struct {
	handlers map[MessageType]func(string)([]byte, error)
}

// NewMessageHandler returns a new MessageHandlers instance to be worked upon
func NewMessageHandler() *MessageHandlers {
	return &MessageHandlers{handlers: make(map[MessageType]func(string)([]byte, error))}
}

// RegisterHandler registers a new message handler
func (mh *MessageHandlers) RegisterHandler(mType MessageType, handler func(string)([]byte, error)) error {
	if _, ok := mh.handlers[mType]; ok {
		return fmt.Errorf("Message handler already exists for the provided type")
	}
	mh.handlers[mType] = handler
	return nil
}

// UnregisterHandler unregisters a registered message type handler
func (mh *MessageHandlers) UnregisterHandler(mType MessageType) error {
	if _, ok := mh.handlers[mType]; ok {
		delete(mh.handlers, mType)
		return nil
	}
	return fmt.Errorf("No message handler registered for the provided message type")
}

// GetHandler returns the message handler for the provided message type
func (mh MessageHandlers) GetHandler(mType MessageType) (func(string)([]byte, error), error) {
	if handler, ok := mh.handlers[mType]; ok {
		return handler, nil
	}
	return nil, fmt.Errorf("Unable to find a handler for the specified message type")
}