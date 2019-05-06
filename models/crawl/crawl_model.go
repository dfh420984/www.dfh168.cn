package crawl

import (
	"regexp"
)

type CrawlModel struct {
}

func (this *CrawlModel) GetContent(content string, rule string) string {
	if content == "" {
		return ""
	}
	reg := regexp.MustCompile(rule)
	result := reg.FindAllStringSubmatch(content, -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

func (this *CrawlModel) GetUrls(content string, rule string) []string {
	reg := regexp.MustCompile(rule)
	result := reg.FindAllStringSubmatch(content, -1)

	var postSets []string
	for _, v := range result {
		postSets = append(postSets, v[1])
	}
	return postSets
}
