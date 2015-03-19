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
	HelloResponse
	GoodbyeResponse
	WorldInfoResponse
	PlayerJoinResponse
	AllPlayersResponse
	AllPlayersPositionResponse
	SendMsgResponse
	ReceiveMsgResponse
)

const Version = "d19032.1"

var AllFrameIds = []FrameId{
	HelloRequest,
	GoodbyeRequest,
	WorldInfoRequest,
	PlayerJoinRequest,
	AllPlayersRequest,
	AllPlayersPositionRequest,
	SendMsgRequest,
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
	Id FrameId
	Short string
	Long string
}

var FrameRequestAliases = map[FrameId]FrameAlias {
	HelloRequest: FrameAlias{
	Long:
		"Say hello to a server",
	Short:
		"HelloRequest",
	},
    GoodbyeRequest: FrameAlias{
	Long:
		"Say goodbye to the server",
	Short:
		"GoodbyeRequest",
	},
	PlayerJoinRequest: FrameAlias{
	Long:
		"Join the world",
	Short:
		"PlayerJoinRequest",
	},
	AllPlayersRequest: FrameAlias{
	Long:
		"List all connected players",
	Short:
		"AllPlayersRequest",
	},
	AllPlayersPositionRequest: FrameAlias{
	Long:
		"List the position of all players",
	Short:
		"AllPlayersPositionRequest",
	},
	WorldInfoRequest: FrameAlias{
	Long:
		"Return metadata about the current world being played",
	Short:
		"WorldInfoRequest",
	},
	SendMsgRequest: FrameAlias{
	Long:
		"Send a message to a player",
	Short:
		"SendMsgRequest",
	},
}

var FrameResponseAliases = map[FrameId]FrameAlias {
	HelloResponse: FrameAlias{
	Long:
		"Say hello to a server",
	Short:
		"HelloResponse",
	},
    GoodbyeResponse: FrameAlias{
	Long:
		"Say goodbye to the server",
	Short:
		"GoodbyeResponse",
	},
	PlayerJoinResponse: FrameAlias{
	Long:
		"Join the world",
	Short:
		"PlayerJoinResponse",
	},
	AllPlayersResponse: FrameAlias{
	Long:
		"List all connected players",
	Short:
		"AllPlayersResponse",
	},
	AllPlayersPositionResponse: FrameAlias{
	Long:
		"List the position of all players",
	Short:
		"AllPlayersPositionResponse",
	},
	WorldInfoResponse: FrameAlias{
	Long:
		"Return metadata about the current world being played",
	Short:
		"WorldInfoResponse",
	},
	SendMsgResponse: FrameAlias{
	Long:
		"Send a message to a player",
	Short:
		"SendMsgResponse",
	},
	ReceiveMsgResponse: FrameAlias{
	Long:
		"Receive a message",
	Short:
		"ReceiveMsgResponse",
	},
}

type HelloRequestFrame struct {
	Version string `codec:"version"`
}

type HelloResponseFrame struct {
	Version string `codec:"version"`
	PlayersOnline uint8 `codec:"online"`
	PlayerCapacity uint8 `codec:"capacity"`
}

func init() {
	for id, alias := range FrameResponseAliases {
		alias.Id = id
	}

	for id, alias := range FrameRequestAliases {
		alias.Id = id
	}
}

