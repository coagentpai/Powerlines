package main

import (
	"net"
	"os"
	"fmt"
	"flag"
	"strconv"
	"io"
	"encoding/hex"
	"code.google.com/p/go-uuid/uuid"
	"github.com/ugorji/go/codec"
)

const (
	CONN_HOST = ""
	CONN_PORT = 3333
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

type ClientPacket struct {
	FrameId int64 `codec:"c"`
	Value interface{} `codec:"v"`
}

type Packet struct {
	command FrameId
	id uuid.UUID
	value interface{}
}

func (p *Player) streamDisconnect() {
	(*p.stream.netConn).Close()
}

func (p *Player) readPacket() (Packet, error) {
	var rawPacket ClientPacket
	err := p.stream.decoder.Decode(&rawPacket)
	if err == io.EOF {
		fmt.Printf("Player %s has disconnected!\n", p.id)
		p.streamDisconnect()
		return Packet{id: p.id}, io.EOF
	}
	fmt.Println(rawPacket)
	decodedPacket := Packet{
		//command: Command(rawPacket.Command),
		value: rawPacket.Value,
		id: p.id,
	}
	return decodedPacket, nil
}

func (p *Player) drainWrites() {
	for {
		bytes := <- *p.stream.writes
		fmt.Printf("Sending to player %s\n", p.id)
		fmt.Println(hex.Dump(bytes))
		conn := *p.stream.netConn
		conn.Write(bytes)
	}
}

func (packet *ClientPacket) Bytes() ([]byte, error) {
	var encodedBytes []byte
	var mh codec.MsgpackHandle
	enc := codec.NewEncoderBytes(&encodedBytes, &mh)
	err := enc.Encode(*packet)
	return encodedBytes, err
}

func (p *Player) newPlayer() error {
	p.id = uuid.NewRandom()
	go p.drainWrites()
	for {
		_, err := p.readPacket()
		if err != nil {
			fmt.Println(err)
			return err
		}
//		commandAlias := CommandAliases[packet.command]
		// fmt.Printf(
		// 	"Command: %s for player %s\n",
		// 	commandAlias.short,
		// 	packet.id.String(),
		// )
		response := ClientPacket{
			//Command: int64(packet.command),
			Value: "Loud and clear!",
		}
		bytes, err := response.Bytes()
		*p.stream.writes <- bytes
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
	writeChan := make(chan []byte, 10)
	newPlayer.stream.writes = &writeChan
	newPlayer.newPlayer()
}

func main() {
	var port int
	flag.IntVar(&port, "p", CONN_PORT, "Port to listen on")
	flag.Parse()

	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST + ":" +  strconv.Itoa(port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + strconv.Itoa(port))
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
