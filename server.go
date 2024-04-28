package lark

import (
	"fmt"
	ByteReader "lark/reader"
	"lark/request"
	"net"
	"strings"
)

type Server struct {
	Address string

	MaxBytesToRead int

	Network string
}

func (server *Server) ListenAndServe() error {
	address := server.Address
	network := server.Network

	if strings.Trim(server.Address, " ") == "" {
		address = ":http"
	}

	if strings.Trim(network, " ") == "" {
		network = TCP
	}

	listener, err := net.Listen(network, address)
	if err != nil {
		fmt.Printf(
			"Encountered error attempting to listen to address: %s using %s",
			address,
			network,
		)
		fmt.Print(err)
		return err
	}

	fmt.Printf(
		"Listening on address: %s using %s",
		address,
		network,
	)

	return server.Serve(listener)
}

func (server *Server) Serve(listener net.Listener) error {
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Print(err)
			return err
		}

		go server.serveConnection(&connection)
	}
}

func (server *Server) serveConnection(connection *net.Conn) {
	defer (*connection).Close()

	byteBuffer := make([]byte, 4069)

	n, err := (*connection).Read(byteBuffer)
	if err != nil {
		fmt.Print(err)
	}

	byteBuffer = byteBuffer[:n]

	fmt.Printf("\nCap: %d Len: %d", cap(byteBuffer), len(byteBuffer))

	if cap(byteBuffer) == len(byteBuffer) {
		fmt.Print("\nCapacity has been met\n")
	}

	byteReader := ByteReader.ByteReader{
		Buffer:   byteBuffer,
		Position: 0,
	}

	request := new(request.Request)

	err = request.HydrateFromByteReader(&byteReader)

	fmt.Printf("%+v\n", request)
}
