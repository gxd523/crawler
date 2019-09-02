package config

const (
	ElasticIndex = "dating_userinfo"

	ItemSaverRpc = "ItemSaverService.Save"
	WorkerRpc    = "CrawlService.Process"

	ParseCityList = "ParseCityList"
	ParseCity     = "ParseCity"
	ParseUserInfo = "ParseUserInfo"

	// Rate limit
	Qps = 5
)
