package protocol

type ContainerId uint16

const (
	HelloRequest ContainerId = iota
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

var AllContainerIds = []ContainerId{
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

type ContainerAlias struct {
	Id ContainerId
	Short string
	Long string
}

var ContainerRequestAliases = map[ContainerId]ContainerAlias {
	HelloRequest: ContainerAlias{
	Long:
		"Say hello to a server",
	Short:
		"HelloRequest",
	},
    GoodbyeRequest: ContainerAlias{
	Long:
		"Say goodbye to the server",
	Short:
		"GoodbyeRequest",
	},
	PlayerJoinRequest: ContainerAlias{
	Long:
		"Join the world",
	Short:
		"PlayerJoinRequest",
	},
	AllPlayersRequest: ContainerAlias{
	Long:
		"List all connected players",
	Short:
		"AllPlayersRequest",
	},
	AllPlayersPositionRequest: ContainerAlias{
	Long:
		"List the position of all players",
	Short:
		"AllPlayersPositionRequest",
	},
	WorldInfoRequest: ContainerAlias{
	Long:
		"Return metadata about the current world being played",
	Short:
		"WorldInfoRequest",
	},
	SendMsgRequest: ContainerAlias{
	Long:
		"Send a message to a player",
	Short:
		"SendMsgRequest",
	},
}

var ContainerResponseAliases = map[ContainerId]ContainerAlias {
	HelloResponse: ContainerAlias{
	Long:
		"Say hello to a server",
	Short:
		"HelloResponse",
	},
    GoodbyeResponse: ContainerAlias{
	Long:
		"Say goodbye to the server",
	Short:
		"GoodbyeResponse",
	},
	PlayerJoinResponse: ContainerAlias{
	Long:
		"Join the world",
	Short:
		"PlayerJoinResponse",
	},
	AllPlayersResponse: ContainerAlias{
	Long:
		"List all connected players",
	Short:
		"AllPlayersResponse",
	},
	AllPlayersPositionResponse: ContainerAlias{
	Long:
		"List the position of all players",
	Short:
		"AllPlayersPositionResponse",
	},
	WorldInfoResponse: ContainerAlias{
	Long:
		"Return metadata about the current world being played",
	Short:
		"WorldInfoResponse",
	},
	SendMsgResponse: ContainerAlias{
	Long:
		"Send a message to a player",
	Short:
		"SendMsgResponse",
	},
	ReceiveMsgResponse: ContainerAlias{
	Long:
		"Receive a message",
	Short:
		"ReceiveMsgResponse",
	},
}

type HelloRequestContainer struct {
	Version string `frame:"version"`
}

type HelloResponseContainer struct {
	Version string `frame:"version"`
	PlayersOnline uint16 `frame:"online"`
	PlayerCapacity uint16 `frame:"capacity"`
}

var ContainerStructMap = map[ContainerId]interface{}{
	HelloResponse: HelloResponseContainer{},
	HelloRequest: HelloRequestContainer{},
}


func init() {
	for id, alias := range ContainerResponseAliases {
		alias.Id = id
	}

	for id, alias := range ContainerRequestAliases {
		alias.Id = id
	}
}

