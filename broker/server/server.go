// Copyrights 2018 Saurabh Badhwar. All Rights Reserved.
// The use of this code is goverened by MIT License
// which can be found in the LICENSE file.

// Package server for handling the UDP server connection to broker.
package server

import (
	"infion/common/messages"
	"strconv"
	"fmt"
	"net"
)

// Constants
const (
	// The buffer size to be used for the message buffer
	BUFFERSIZE = 32768

	// Error severity related constants
	FATAL = "FATAL"
	WARNING = "WARNING"
	ERROR = "ERROR"

	// Connection Error Related Error Numbers
	ADDRESSERROR = 100
	CONNECTIONFAILED = 101
	CONNECTIONCLOSED = 102
	PARTIALDATA = 103
)

// ConnectionError defines the structure for providing information
// related to the errors that happen during the establishment of
// connection for the server or the client.
type ConnectionError struct {
	errNo		int
	errMsg		string
	severity	string
}

// NewConnectionError returns a new error related to the connection
func NewConnectionError(errNo int, errMsg string, severity string) *ConnectionError {
	return &ConnectionError{errNo: errNo, errMsg: errMsg, severity: severity}
}

func (connError *ConnectionError) Error() string {
	return fmt.Sprintf("[%s]%d: %s", connError.severity, connError.errNo, connError.errMsg)
}

// Server specifies the structure that needs to be used for establishing
// a new UDP based server for the broker.
type Server struct {
	// udpAddr is the UDP translated address by the net package
	// to be used by the connection library
	udpAddr		*net.UDPAddr

	// udpConn is the object instance that is generated after the
	// call to the connection function.
	udpConn		*net.UDPConn

	// clientPool holds the ClientPool object 
	clientPool	*ClientPool

	// messageHandlers defines a new message handling pool
	messageHandlers *messages.MessageHandlers
}

// NewServer establishes a new UDP server to listen to the incoming
// connections and handle the message relay
func NewServer(host string, port int) (*Server, error) {
	var err error
	uri := host + ":" + strconv.Itoa(port)
	server := new(Server)

	server.udpAddr, err = net.ResolveUDPAddr("udp", uri)
	if err != nil {
		fmt.Println("Error translating UDP address")
		return nil, NewConnectionError(ADDRESSERROR, "Unable to resolve UDP address", ERROR)
	}

	server.clientPool = NewClientPool()
	server.messageHandlers = messages.NewMessageHandler()	
	return server,nil
}


// Listen makes the server to accept new connections as they arrive
// while also maintaining a queue of the address units that are connecting to
// the server.
// The method does not process the incoming messages and as soon as they arrive
// the message is handed over to the message request handler that runs in the
// concurrent thread, hence allowing our Listen method to work on the new
// incoming message requests.
func (s *Server) Listen() error {
	buf := make([]byte, BUFFERSIZE)
	var bytesRead int
	var clientAddr *net.UDPAddr
	var err error

	s.udpConn, err = net.ListenUDP("udp", s.udpAddr)
	if err != nil {
		fmt.Println("Unable to start the UDP Server")
		return NewConnectionError(CONNECTIONFAILED, "Unable to bind to the address for listening", FATAL)
	}
	defer s.udpConn.Close()

	// Start listening to the incoming messages
	for {
		bytesRead, clientAddr, err = s.udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error occured while trying to read data")
		}
		fmt.Println("Read ", strconv.Itoa(bytesRead), "from ", clientAddr)
	}

	return nil
}

