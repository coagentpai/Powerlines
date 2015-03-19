package main

type FrameId uint16

const (
	HelloRequest FrameId = iota
	GoodbyeRequest
	WorldInfoRequest
	PlayerJoinRequest
	AllPlayersRequest
	AllPlayersPositionRequest
	SendMsgRequest
	ReceiveMsgRequest
	HelloResponse
	GoodbyeResponse
	WorldInfoResponse
	PlayerJoinResponse
	AllPlayersResponse
	AllPlayersPositionResponse
	SendMsgResponse
	ReceiveMsgResponse
)
var AllFrameIds = []FrameId{
	HelloRequest,
	GoodbyeRequest,
	WorldInfoRequest,
	PlayerJoinRequest,
	AllPlayersRequest,
	AllPlayersPositionRequest,
	SendMsgRequest,
	ReceiveMsgRequest,
	HelloResponse,
	GoodbyeResponse,
	WorldInfoResponse,
	PlayerJoinResponse,
	AllPlayersResponse,
	AllPlayersPositionResponse,
	SendMsgResponse,
	ReceiveMsgResponse,
}

type FrameAlias struct {
	id FrameId
	short string
	long string
}

var FrameRequestAliases = map[FrameId]FrameAlias {
	HelloRequest: FrameAlias{
	long:
		"Say hello to a server",
	short:
		"HelloRequest",
	},
    GoodbyeRequest: FrameAlias{
	long:
		"Say goodbye to the server",
	short:
		"GoodbyeRequest",
	},
	PlayerJoinRequest: FrameAlias{
	long:
		"Join the world",
	short:
		"PlayerJoinRequest",
	},
	AllPlayersRequest: FrameAlias{
	long:
		"List all connected players",
	short:
		"AllPlayersRequest",
	},
	AllPlayersPositionRequest: FrameAlias{
	long:
		"List the position of all players",
	short:
		"AllPlayersPositionRequest",
	},
	WorldInfoRequest: FrameAlias{
	long:
		"Return metadata about the current world being played",
	short:
		"WorldInfoRequest",
	},
	SendMsgRequest: FrameAlias{
	long:
		"Send a message to a player",
	short:
		"SendMsgRequest",
	},
	ReceiveMsgRequest: FrameAlias{
	long:
		"Receive a message",
	short:
		"ReceiveMsgRequest",
	},
}

var FrameAliases = map[FrameId]FrameAlias {}

func init() {
	for id, alias := range FrameAliases {
		alias.id = id
	}
}

