package config

const (
	ElasticIndex = "dating_userinfo"

	ItemSaverRpc = "ItemSaverService.Save"
	WorkerRpc    = "CrawlService.Process"
	DuplicateRpc    = "RepeatService.IsDuplicateUrl"

	ParseCityList = "ParseCityList"
	ParseCity     = "ParseCity"
	ParseUserInfo = "ParseUserInfo"

	// Rate limit
	Qps = 5
)
