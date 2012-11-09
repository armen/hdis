// Very memory efficient plain key-value store on top of Redis
// As it's described in http://redis.io/topics/memory-optimization

package hdis

import (
	"github.com/garyburd/redigo/redis"

	"errors"
	"strings"
)

type Conn struct {
	Conn redis.Conn
}

var NotAHashCommandError = errors.New("Given command is not a hash command")

func getKeyField(bigkey string) (key, field string) {

	parts := strings.SplitN(bigkey, ":", 2)

	if len(parts) != 2 {
		return bigkey, field
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
	return c.Do("HGET", bigkey)
}

func (c Conn) Set(bigkey string, value interface{}) (interface{}, error) {
	return c.Do("HSET", bigkey, value)
}

func (c Conn) Do(commandName string, bigkey string, args ...interface{}) (interface{}, error) {

	if !strings.HasPrefix(commandName, "H") {
		return nil, NotAHashCommandError
	}

	key, field := getKeyField(bigkey)

	switch strings.ToUpper(commandName) {
	case "HGETALL", "HKEYS", "HLEN", "HVALS":
		arguments := append([]interface{}{key}, args...)
		return c.Conn.Do(commandName, arguments...)
	}

	arguments := append([]interface{}{key, field}, args...)
	return c.Conn.Do(commandName, arguments...)
}
