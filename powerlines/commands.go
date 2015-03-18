package main

type Command uint16

const (
	Hello Command = iota
	Goodbye
	WorldInfo
	PlayerJoin
	AllPlayers
	AllPlayersPosition
	SendMsg
	RecieveMsg
)
var AllCommands = []Command{
	Hello,
	Goodbye,
	WorldInfo,
	PlayerJoin,
	AllPlayers,
	AllPlayersPosition,
	SendMsg,
	RecieveMsg,
}

type CommandAlias struct {
	command Command
	short string
	long string
}

var CommandAliases = map[Command]CommandAlias {
	Hello: CommandAlias{
	long:
		"Say hello to a server",
	short:
		"Hello",
	},
    Goodbye: CommandAlias{
	long:
		"Say goodbye to the server",
	short:
		"Goodbye",
	},
	PlayerJoin: CommandAlias{
	long:
		"Join the world",
	short:
		"PlayerJoin",
	},
	AllPlayers: CommandAlias{
	long:
		"List all connected players",
	short:
		"AllPlayers",
	},
	AllPlayersPosition: CommandAlias{
	long:
		"List the position of all players",
	short:
		"AllPlayersPosition",
	},
	WorldInfo: CommandAlias{
	long:
		"Return metadata about the current world being played",
	short:
		"WorldInfo",
	},
}

func init() {
	for command, alias := range CommandAliases {
		alias.command = command
	}
}

