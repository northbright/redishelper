package redishelper

import (
	"fmt"
	"strings"

	"github.com/garyburd/redigo/redis"
)

// GetInfo returns the Redis Info.
//
//     Return:
//         info: key: section name, value: map[string]string. key: info name, value: info value.
func GetInfo(c redis.Conn) (info map[string]map[string]string, err error) {
	info = make(map[string]map[string]string)
	sections := []string{}

	ret, err := redis.String(c.Do("INFO"))
	if err != nil {
		goto end
	}

	ret = strings.Trim(ret, "\r\n")
	sections = strings.Split(ret, "\r\n\r\n")

	for _, s := range sections {
		items := strings.Split(s, "\r\n")
		l := len(items)
		if l == 0 {
			err = fmt.Errorf("No INFO sections found.")
			goto end
		}

		arr := strings.Split(items[0], "# ")
		if len(arr) != 2 {
			err = fmt.Errorf("No section name found.")
			goto end
		}

		section := arr[1]
		info[section] = make(map[string]string)
		for i := 1; i < l; i++ {
			arr := strings.Split(items[i], ":")
			if len(arr) != 2 {
				err = fmt.Errorf("No parameter found.")
				goto end
			}
			k, v := arr[0], arr[1]
			info[section][k] = v
		}
	}

end:
	return info, err
}
