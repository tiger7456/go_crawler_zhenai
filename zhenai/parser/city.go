package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
const cityUrlRe = `href="(http://www.zhenai.com/zhenghun/[^"]+)"`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllStringSubmatch(string(contents), -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		url := m[1]
		result.Requests = append(
			result.Requests, engine.Request{
				Url: m[1],
				ParserFunc: func(content []byte) engine.ParseResult {
					return ParseProfile(content, url)
				},
			})
	}
	re = regexp.MustCompile(cityUrlRe)
	matches = re.FindAllStringSubmatch(string(contents), -1)
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        m[1],
				ParserFunc: ParseCity,
			})
	}
	return result
}
