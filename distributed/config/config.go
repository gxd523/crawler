package config

const (
	ElasticIndex = "dating_userinfo"

	ItemSaverPort = 1234
	WorkerPort    = 1235

	ItemSaverRpc = "ItemSaverService.Save"
	WorkerRpc    = "CrawlService.Process"

	ParseCityList = "ParseCityList"
	ParseCity     = "ParseCity"
	ParseUserInfo = "ParseUserInfo"
)
