package lib

import (
	"strings"
)

func GetURISuffix(uri string) string {
	if len(uri) == 0 {
		return ""
	}
	i := strings.Index(uri[1:], "/")
	if i == -1 {
		// URL does not contain suffix
		return ""
	}
	return uri[i+1:]
}