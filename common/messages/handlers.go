// Copyrights 2018 Saurabh Badhwar. All Rights Reserved.
// The use of this code is governed by MIT License
// which can be found in the LICENSE file.

package messages

import (
	"fmt"
)

// MessageHandler defines the type for function signature to be used by the
// message handlers.
type MessageHandler func(string)([]byte, error)

// MessageHandlers defines a type to register different message
// handlers respective of the types they are supposed to handle.
// For a new message handler, they have to provide the type of message
// they can handle and the handler function that should be called
// with the payload.
type MessageHandlers struct {
	handlers map[MessageType]MessageHandler
}

// NewMessageHandler returns a new MessageHandlers instance to be worked upon
func NewMessageHandler() *MessageHandlers {
	return &MessageHandlers{handlers: make(map[MessageType]MessageHandler)}
}

// RegisterHandler registers a new message handler
func (mh *MessageHandlers) RegisterHandler(mType MessageType, handler MessageHandler) error {
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
func (mh MessageHandlers) GetHandler(mType MessageType) (MessageHandler, error) {
	if handler, ok := mh.handlers[mType]; ok {
		return handler, nil
	}
	return nil, fmt.Errorf("Unable to find a handler for the specified message type")
}

// GenericHandler defines a generic way of dealing with the incoming message
// payloads.
func GenericHandler(payload string) ([]byte, error) {
	fmt.Println(payload)
	return []byte(payload), nil
}