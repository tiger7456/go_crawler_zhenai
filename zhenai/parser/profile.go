package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"

	"github.com/tidwall/gjson"
)

const profileRe = `__INITIAL_STATE__=(.+);\(function`

func ParseProfile(content []byte) engine.ParseResult {
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

	result := engine.ParseResult{
		Items:    []interface{}{profile},
		Requests: nil,
	}
	return result
}
