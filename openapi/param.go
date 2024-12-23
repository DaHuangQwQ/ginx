package openapi

import (
	"regexp"
)

// parsePathParams gives the list of path parameters in a path.
// Example : /item/:user/:id -> [user, id]
func parsePathParams(path string) []string {
	// 正则表达式匹配以 ":" 开头的参数
	re := regexp.MustCompile(`:([a-zA-Z0-9_]+)`)

	matches := re.FindAllStringSubmatch(path, -1)

	var params []string
	for _, match := range matches {
		if len(match) > 1 {
			params = append(params, match[1])
		}
	}
	return params
}
