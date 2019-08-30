package parser

import (
	"crawler/distributed/config"
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
	"strings"
)

const (
	// "basicInfo":["未婚","30岁","射手座(11.22-12.21)","160cm","43kg","工作地:上海虹口区","月收入:1.2-2万","金融","硕士"],
	// "detailInfo":["汉族","籍贯:上海","体型:运动员型","稍微抽一点烟","不喝酒","和家人同住","未买车","没有小孩","是否想要孩子:不想要孩子","何时结婚:时机成熟就结婚"]
	// "educationString": "大专"
	userInfoRegex = `"basicInfo"[^[]+([^]]+])[^[]+([^]]+])[\d\D]+?"educationString"[^"]+"([^"]+)`

	// http://album.zhenai.com/u/1599085703
	idRegex = `http://album.zhenai.com/u/([\d]+)`
)

var (
	userInfoCompile    = regexp.MustCompile(userInfoRegex)
	idCompile          = regexp.MustCompile(idRegex)
	incomeCompile      = regexp.MustCompile(`月收入:([^"]+)`)
	weightCompile      = regexp.MustCompile(`([\d]{2})kg`)
	ageCompile         = regexp.MustCompile(`"([\d]{2})岁"`)
	heightCompile      = regexp.MustCompile(`"([\d]{3})cm"`)
	xinzuoCompile      = regexp.MustCompile(`([\p{Han}]{2}座)\([\d-.]+\)`)
	cityCompile        = regexp.MustCompile(`"工作地:([^"]+)"`)
	birthplaceCompile  = regexp.MustCompile(`"籍贯:([^"]+)"`)
	nationalityCompile = regexp.MustCompile(`"([\p{Han}]+族)"`)
	wannaChildCompile  = regexp.MustCompile(`"是否想要孩子:([^"]+)"`)
)

func parseUserInfo(bytes []byte, username string, gender string, url string) *engine.ParseResult {
	userInfoSubmatch := userInfoCompile.FindSubmatch(bytes)

	userInfo := model.UserInfo{}
	userInfo.Name = username
	userInfo.Gender = gender

	if userInfoSubmatch != nil {
		for k, v := range userInfoSubmatch {
			switch k {
			case 1:
				userInfo.Marriage = deserializeSlice(v)[0]
				userInfo.Age = stringToInt(getSubMatch(ageCompile, v))
				userInfo.Xinzuo = getSubMatch(xinzuoCompile, v)
				userInfo.Height = stringToInt(getSubMatch(heightCompile, v))
				userInfo.Weight = stringToInt(getSubMatch(weightCompile, v))
				userInfo.City = getSubMatch(cityCompile, v)
				userInfo.Income = getSubMatch(incomeCompile, v)
				break
			case 2:
				userInfo.Nationality = getSubMatch(nationalityCompile, v)
				userInfo.Birthplace = getSubMatch(birthplaceCompile, v)
				userInfo.House = getMatchHouse(v)
				userInfo.Car = getMatchCar(v)
				userInfo.WannaChild = getSubMatch(wannaChildCompile, v)
				userInfo.HaveChild = getMatchChild(v)
				break
			case 3:
				userInfo.Education = string(v)
				break
			}
		}
	}
	item := engine.Item{
		Type:    "zhenai",
		Id:      idCompile.FindStringSubmatch(url)[1],
		Url:     url,
		Payload: userInfo,
	}
	return &engine.ParseResult{Item: &item}
}

func deserializeSlice(s []byte) []string {
	trimString := strings.ReplaceAll(string(s), "\"", "")
	trimString = strings.ReplaceAll(trimString, "[", "")
	trimString = strings.ReplaceAll(trimString, "]", "")
	trimString = strings.ReplaceAll(trimString, " ", "")
	return strings.Split(trimString, ",")
}

func stringToInt(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return 0
}

func getSubMatch(reg *regexp.Regexp, bytes []byte) string {
	if submatch := reg.FindSubmatch(bytes); len(submatch) > 0 {
		return string(submatch[1])
	}
	return "无"
}

func getMatchHouse(bytes []byte) string {
	s := string(bytes)
	switch {
	case strings.Contains(s, "已购房"):
		return "已购房"
	case strings.Contains(s, "租房"):
		return "租房"
	case strings.Contains(s, "住在单位宿舍"):
		return "住在单位宿舍"
	case strings.Contains(s, "和家人同住"):
		return "和家人同住"
	case strings.Contains(s, "打算婚后购房"):
		return "打算婚后购房"
	default:
		//log.Printf("未找到购房信息....%s\n", bytes)
		return "未填写"
	}
}

func getMatchCar(bytes []byte) string {
	s := string(bytes)
	switch {
	case strings.Contains(s, "未买车"):
		return "未买车"
	case strings.Contains(s, "已买车"):
		return "已买车"
	default:
		//log.Printf("未找到购车信息....%s\n", bytes)
		return "未填写"
	}
}

func getMatchChild(bytes []byte) string {
	s := string(bytes)
	switch {
	case strings.Contains(s, "没有小孩"):
		return "没有小孩"
	case strings.Contains(s, "有孩子但不在身边"):
		return "有孩子但不在身边"
	case strings.Contains(s, "有孩子且偶尔会一起住"):
		return "有孩子且偶尔会一起住"
	case strings.Contains(s, "有孩子且住在一起"):
		return "有孩子且住在一起"
	default:
		//log.Printf("未找到子女信息....%s\n", bytes)
		return "未填写"
	}
}

type UserInfoParser struct {
	Username string
	Gender   string
}

func (p *UserInfoParser) Parse(bytes []byte, url string) *engine.ParseResult {
	return parseUserInfo(bytes, p.Username, p.Gender, url)
}

func (p *UserInfoParser) Serialize() (funcName string, args interface{}) {
	return config.ParseUserInfo, [2]string{p.Username, p.Gender}
}
