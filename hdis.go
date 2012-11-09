// Very memory efficient plain key-value store on top of Redis
// As it's described in http://redis.io/topics/memory-optimization

package hdis

import (
	"github.com/garyburd/redigo/redis"

	"strings"
)

type Conn struct {
	redis.Conn
}

func getKeyField(bigkey string) (key, field string) {

	parts := strings.SplitN(bigkey, ":", 2)

	if len(parts) != 2 {
		return
	}

	if len(parts[1]) > 2 {
		key = parts[0] + ":" + parts[1][0:2]
		field = parts[1][2:]
	} else {
		key = parts[0] + ":"
		field = parts[1]
	}

	return
}

func (c Conn) Get(bigkey string) (interface{}, error) {
	key, field := getKeyField(bigkey)
	return c.Do("HGET", key, field)
}

func (c Conn) Set(bigkey string, value interface{}) (interface{}, error) {
	key, field := getKeyField(bigkey)
	return c.Do("HSET", key, field, value)
}
