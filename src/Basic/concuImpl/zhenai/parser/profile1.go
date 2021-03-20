package parser

import (
	"goLangLearn/concuImpl/engine"
	"goLangLearn/concuImpl/model"
	"regexp"
)

var purpleRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
var pinkRe = regexp.MustCompile(
	`<div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div>`)

func parseProfile1(contents []byte, name string) engine.ParseResult {
	profile := model.Profile1{}
	profile.Name = name

	profile.Marriage = extract(contents, purpleRe, 0)
	profile.Age = extract(contents, purpleRe, 1)
	profile.Xinzuo = extract(contents, purpleRe, 2)
	profile.Height = extract(contents, purpleRe, 3)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extract(contents []byte, re *regexp.Regexp, index int) string {
	match := re.FindAllSubmatch(contents, -1)

	if len(match) >= 4 {
		return string(match[index][1])
	} else {
		return ""
	}
}
