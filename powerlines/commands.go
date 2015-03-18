package main

type Command uint16

const (
	Hello Command = iota
	Goodbye
	WorldInfo
	AllPlayers
	AllPlayersPosition
	SendMsg
	RecieveMsg
)

var CommandReadable = map[Command]string{
	Hello: "Say hello to a server",
    Goodbye: "ahskjahsdj",
	AllPlayers: "sakljsdal",
	AllPlayersPosition: "klsjsshkjsa",
	WorldInfo: "sdakllsdaj",
	
}

