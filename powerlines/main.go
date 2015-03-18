package main

import (
	"net"
	"os"
	"fmt"
	"io"
	"code.google.com/p/go-uuid/uuid"
	"github.com/ugorji/go/codec"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

type PlayerVertex struct {
	x, y, z float64
}

type PacketStream struct {
	writes *chan []byte
	netConn *net.Conn
	decoder *codec.Decoder
}


type Player struct {
	position PlayerVertex
	look PlayerVertex
	name string
	id uuid.UUID
	stream PacketStream
}

type PacketRaw struct {
	Command int64 `codec:"c"`
	Value interface{} `codec:"v"`
}

type Packet struct {
	command uint16
	id uuid.UUID
	value interface{}
}

func (p *Player) streamDisconnect() {
	(*p.stream.netConn).Close()
}

func (p *Player) readPacket() (Packet, error) {
	var rawPacket PacketRaw
	err := p.stream.decoder.Decode(&rawPacket)
	if err == io.EOF {
		fmt.Printf("Player %s has disconnected!\n", p.id)
		p.streamDisconnect()
		return Packet{id: p.id}, io.EOF
	}
	fmt.Println(rawPacket)
	decodedPacket := Packet{
		command: uint16(rawPacket.Command),
		value: rawPacket.Value,
		id: p.id,
	}
	return decodedPacket, nil
}

func (p *Player) newPlayer() error {
	p.id = uuid.NewRandom()
	for {
		packet, err := p.readPacket()
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("command: %d for player %s\n", packet.command, packet.id.String())
	}
	return nil
}

// Handles incoming requests.
func handleConnection(conn net.Conn) {
	var newPlayer Player
	var mh codec.MsgpackHandle
	defer conn.Close()
	newPlayer.stream.netConn = &conn
	newPlayer.stream.decoder = codec.NewDecoder(conn, &mh)
	newPlayer.newPlayer()
}

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleConnection(conn)
	}

}
