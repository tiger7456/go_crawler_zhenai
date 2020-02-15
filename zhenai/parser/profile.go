package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"

	"github.com/tidwall/gjson"
)

const profileRe = `__INITIAL_STATE__=(.+);\(function`
const urlRe = `http://album.zhenai.com/u/([\d]+)`

func ParseProfile(content []byte, url string) engine.ParseResult {
	profile := model.Profile{}
	re := regexp.MustCompile(profileRe)
	match := re.FindStringSubmatch(string(content))
	data := ``
	if len(match) >= 2 {
		data = string(match[1])
	}
	// 解析用户信息
	objectInfo := gjson.Get(data, "objectInfo")
	profile.Name = objectInfo.Get("nickname").String()
	profile.Age = int(objectInfo.Get("age").Int())
	profile.Gender = objectInfo.Get("genderString").String()
	heightStr := objectInfo.Get("heightString").String()
	profile.Height, _ = strconv.Atoi(regexp.MustCompile(`\d+`).FindString(heightStr))
	profile.Education = objectInfo.Get("educationString").String()
	profile.Hukou = objectInfo.Get("workProvinceCityString").String()
	profile.Income = objectInfo.Get("salaryString").String()
	profile.Marriage = objectInfo.Get("marriageString").String()
	profile.Occupation = objectInfo.Get("basicInfo.6").String()
	profile.Xinzuo = objectInfo.Get("basicInfo.2").String()

	re = regexp.MustCompile(urlRe)
	match = re.FindStringSubmatch(string(url))
	var id string
	if len(match) > 1 {
		id = match[1]
	}
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      id,
				Payload: profile,
			},
		},
		Requests: nil,
	}
	return result
}
