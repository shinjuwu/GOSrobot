package lobbyservice

type LSCommandLayer struct {
	Sys      string      `json:"sys"`
	Cmd      string      `json:"cmd"`
	Sn       int         `json:"sn"`
	IsEncode bool        `json:"isEncode"`
	Data     interface{} `json:"data"`
}

type LSRespLayer struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Ret     string      `json:"ret"`
	Data    interface{} `json:"data"`
}

type LSIPGet struct {
	PlatformID    int    `json:"platformId"`
	MemberCode    int    `json:"memberCode"`
	AgentID       int    `json:"agentId"`
	GameID        int    `json:"gameId"`
	GameCode      string `json:"gameCode"`
	Token         string `json:"token"`
	Account       string `json:"account"`
	Password      string `json:"password"`
	ClientType    string `json:"clientType"`
	ClientVersion string `json:"clientVersion"`
	GameTicket    string `json:"gameTicket"`
	ServerIdx     int    `json:"serverIdx"`
}

type IPGetRespData struct {
	ServerID   int    `json:"ServerId`
	IP         string `json:"ip"`
	ServerPort int    `json:"serverPort"`
	GameTicket string `json:"gameTicket"`
	ResUrl     string `json:`
}

type LSserverAddData struct {
	ServerID      int    `json:"serverId"`
	ServerVersion string `json:"serverVersion"`
	ServerMode    string `json:"serverMode"`
	ReleaseMode   int    `json:"releaseMode"`
	PlatformID    int    `json:"platformId"`
	MemberCode    int    `json:"memberCode"`
	AgentID       int    `json:"agentId"`
	IP            string `json:"ip"`
	IP2           string `json:"ip2"`
	DBname        string `json:"dbName"`
	DBport        int    `json:"dbPort"`
	ServerPort    int    `json:"serverPort"`
	HTTPPort      int    `json:"httpPort"`
}

// type serverAddResData struct {
// 	ServerID      int    `json:"serverId"`
// 	ServerVersion string `json:"serverVersion"`
// 	PlatformID    int    `json:"platformId"`
// 	MemberCode    int    `json:"memberCode"`
// 	AgentID int
// }

type LSmemberLogin struct {
	ServerID            int    `json:"serverId"`
	PlatformID          int    `json:"platformId"`
	ThirdPartyUserID    int64  `json:"thirdPartyUserId"`
	ThirdPartyUserIDStr string `json:"thirdPartyUserIdStr"`
	TotalUser           int    `json:"totalUser"`
	Token               string `json:"token"`
	GameTicket          string `json:"gameTicket"`
	Account             string `json:"account"`
	Password            string `json:"password"`
	GameID              int    `json:"gameId"`
	GameCode            string `json:"gameCode"`
	ClientType          string `json:"clientType"`
	ClientVersion       string `json:"clientVersion"`
	IsLogin             bool   `json:"isLogin"`
}

type MemberLoginResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Ret     string      `json:"ret"`
	Data    interface{} `json:"data"`
}

type memberLoginResData struct {
	PlatformID          int    `json:"platformId"`
	MemberCode          int    `json:"memberCode"`
	AgentID             int    `json:"agentId"`
	UserID              int    `json:"user_id"`
	ThirdPartyUserID    int64  `json:"thirdPartyUserId"`
	ThirdPartyUserIDStr string `json:"thirdPartyUserIdStr"`
	Account             string `json:"account"`
	Nickname            string `json:"nickname"`
	GameTicket          string `json:"gameTicket"`
	Token               string `json:"token"`
	Balance             int64  `json:"balance"`
	BalanceCI           int64  `json:"balance_ci"`
	BalanceWin          int64  `json:"balance_win"`
	VIPRank             int    `json:"vip_rank"`
	LoginTime           string `json:"loginTime"`
	CreateTime          string `json:"createTime"`
	UpdateTime          string `json:"updateTime"`
	Status              int    `json:"status"`
	OrderState          int    `json:"orderState"`
}

type LSMemberLogout struct {
	ServerID            int    `json:"serverId"`
	PlatformID          int    `json:"platformId"`
	UserID              int    `json:"user_id"`
	ThirdPartyUserID    int    `json:"thirdPartyUserId"`
	ThirdPartyUserIDStr string `json:"thirdPartyUserIdStr"`
	TotalUser           int    `json:"totalUser"`
	Token               string `json:"token"`
	GameTicket          string `json:"gameTicket"`
	Account             string `json:"account"`
	Password            string `json:"password"`
	GameID              int    `json:"gameId"`
	GameCode            string `json:"gameCode"`
	ClientType          string `json:"clientType"`
	ClientVersion       string `json:"clientVersion"`
	IsLogin             bool   `json:"isLogin"`
}

type MemberLogoutResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Ret     string              `json:"ret"`
	Data    memberLogoutResData `json:"data"`
}

type memberLogoutResData struct {
	PlatformID          int    `json:"platformId"`
	MemberCode          int    `json:"memberCode"`
	AgentID             int    `json:"agentId"`
	UserID              int    `json:"user_id"`
	ThirdPartyUserID    int64  `json:"thirdPartyUserId"`
	ThirdPartyUserIDStr string `json:"thirdPartyUserIdStr"`
	Account             string `json:"account"`
	Nickname            string `json:"nickname"`
	GameTicket          string `json:"gameTicket"`
	Token               string `json:"token"`
	Balance             int64  `json:"balance"`
	BalanceCI           int64  `json:"balance_ci"`
	BalanceWin          int64  `json:"balance_win"`
	VIPRank             int    `json:"vip_rank"`
	LoginTime           string `json:"loginTime"`
	CreateTime          string `json:"createTime"`
	UpdateTime          string `json:"updateTime"`
	Status              int    `json:"status"`
	OrderState          string `json:"orderState"`
}

type LSPlatformInfoGet struct {
	ServerID int `json:"serverId"`
}

type PlatformInfoGetResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Ret     string                 `json:"ret"`
	Data    platformInfoGetResData `json:"data"`
}

type platformInfoGetResData struct {
	PlatformID       int    `json:"platformId"`
	MemberCode       int    `json:"memberCode"`
	AgentID          int    `json:"agentId"`
	PlatformName     string `json:"platformName"`
	PlatformAccount  string `json:"platformAccount"`
	PlatformPassword string `json:"platformPassword"`
	WebapiMode       int    `json:"webapiMode"`
	GSWebapiMode     int    `json:"gsWebapiMode"`
}

type LSGameInfoGet struct {
	ServerID int `json:"serverId"`
}
type GameInfoGetResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Ret     string              `json:"ret"`
	Data    gameInfoGetResponse `json:"data"`
}

type gameInfoGetResponse struct {
	PlatformID          int    `json:"platformId"`
	GameID              int    `json:"gameId"`
	GameName            string `json:"gameName"`
	GameEnName          string `json:"gameEnName"`
	GameMode            int    `json:"gameMode"`
	TableDestoryMode    int    `json:"tableDestoryMode"`
	OpenTableMax        int    `json:"openTableMax"`
	TablePlayerMax      int    `json:"tablePlayerMax"`
	DisconnectCleanData int    `json:"disconnectCleanData"`
	AfterKickBefore     int    `json:"afterKickBefore"`
	BetClusterSecs      int    `json:"betClusterSecs"`
	PlayTimeMax         int    `json:"playTimeMax"`
	InPlayTime          int    `json:"inPlayTime"`
	SettlementTimeMax   int    `json:"settlementTimeMax"`
	Enable              int    `json:"enable"`
}

type LSLobbyInfoGet struct {
	ServerID int `json:"serverId"`
}

type LobbyInfoGetResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Ret     string              `json:"ret"`
	Data    lobbyInfoGetResData `json:"data"`
}

type lobbyInfoGetResData struct {
	PlatformID   int    `json:"platformId"`
	LobbyID      int    `json:"lobbyId"`
	GameID       int    `json:"gameId"`
	LobbyName    string `json:"lobbyName"`
	LobbyMatchID int    `json:"lobbyMatchId"`
	TotalWater1  int64  `json:"total_water1"`
	TotalWater2  int64  `json:"total_water2"`
	BetLevel     []int  `json:"betLevel"`
	Hot          int    `json:"hot"`
}

type LSGameDataGet struct {
	ServerID   int    `json:"serverId"`
	PlatformID int    `json:"platformId"`
	LobbyID    int    `json:"lobbyId"`
	GameID     int    `json:"gameId"`
	Account    string `json:"account"`
	UserID     int    `json:"user_id"`
}

type GameDataGetResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Ret     string             `json:"ret"`
	Data    GameDataGetResData `json:"data"`
}

//TODO自訂資料的回傳值格式
type GameDataGetResData struct {
	Type        string  `json:"type"`
	BetID       int64   `json:"betID"`
	UserID      string  `json:"userID"`
	Index       int     `json:"index"`
	Record      string  `json:"record"`
	BetMultiple float64 `json:"betMultiple"`
	BetKey      int     `json:"betKey"`
	Cheater     int     `json:"cheater"`
	LifePoint   int     `josn:"lifePoint"`
}

type LSGameDataSet struct {
	ServerID         int             `json:"serverId"`
	PlatformID       int             `json:"platformId"`
	MemberCode       int             `json:"memberCode"`
	AgentID          int             `json:"agentId"`
	LobbyID          int             `json:"lobbyId"`
	GameID           int             `json:"gameId"`
	GameCode         string          `json:"gameCode"`
	Account          string          `json:"account"`
	UserID           int             `json:"user_id"`
	ThirdPartyUserID int             `json:"thirdPartyUserId"`
	GameData         GameDataSetData `json:"gameData"`
}

//TODO:確認設定資料的資料
type GameDataSetData struct {
	Type        string  `json:"type"`
	BetID       uint64  `json:"betID"`
	UserID      string  `json:"userID"`
	Index       int     `json:"index"`
	Record      string  `json:"record"`
	BetMultiple float64 `json:"betMultiple"`
	BetKey      int     `json:"betKey"`
	Cheater     int     `json:"cheater"`
	LifePoint   int     `josn:"lifePoint"`
}

type GameDataSetResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ret     string `json:"ret"`
	Data    string `json:"data"`
}

type gameDataSetResData struct {
	ServerID         int    `json:"serverId"`
	PlatformID       int    `json:"platformId"`
	MemberCode       int    `json:"memberCode"`
	AgentID          int    `json:"agentId"`
	Account          string `json:"account"`
	UserID           int    `json:"user_id"`
	ThirdPartyUserID int    `josn:"thirdPartyUserId"`
	IsSuccess        bool   `json:"isSuccess"`
}
