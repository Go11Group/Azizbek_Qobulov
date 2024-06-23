package postgres

import (
	"fmt"
	"strings"
)

// ReplaceQueryParams SQL so'rovidagi nomlangan parametrlarni tegishli to'ldiruvchilar bilan almashtiradi
// va o'zgartirilgan so'rov va argumentlar ro'yxatini qaytaradi.
func ReplaceQueryParams(query string, params map[string]interface{}) (string, []interface{}) {
	args := make([]interface{}, 0, len(params))
	for key, val := range params {
		query = strings.ReplaceAll(query, ":"+key, fmt.Sprintf("$%d", len(args)+1))
		args = append(args, val)
	}
	return query, args
}
