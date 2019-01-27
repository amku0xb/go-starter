package constant

const (
	//app const
	ServerStarted    = "############################## SERVER STARTED ##############################"
	InitLogger       = "INIT LOGGER SUCCESS"
	InitNewRelic     = "INIT NEW RELIC SUCCESS"
	InitNewRelicFail = "INIT NEW RELIC FAILURE"
	InitMgoDbSuccess = "MGO DB INIT SUCCESS"
	MgoUrlParseErr   = "MGO DB URL PARSE FAILURE"
	MgoSesInitFail   = "INIT MGO SESSION FAILURE"
	InitMongoSuccess = "MONGO DB INIT SUCCESS : %v"
	InitMongoFail    = "MONGO DB INIT FAILURE : %v"
	InitRedisSuccess = "INIT REDIS SUCCESS"
	InitRedisFail    = "INIT REDIS FAILURE"

	//req methods
	Post            = "POST"
	Get             = "GET"
	ContentTypeKey  = "Content-Type"
	JsonContentType = "application/json"

	//errors
	NewRelicAgentNotFound = "new relic agent not found"
	MsgParseFail          = "failed to parse the message"
	ExtApiCallFail        = "failed to receive response in JSONHTTPPost"
	FailToUpdateLock      = "fail to update lock"
	FailToFindLock        = "fail to find lock by criteria"
)
