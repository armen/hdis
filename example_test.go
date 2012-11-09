package hdis_test

import (
	"github.com/armen/hdis"
	"github.com/garyburd/redigo/redis"

	"fmt"
)

func ExampleGet() {

	c, _ := redis.Dial("tcp", ":6379")
	defer c.Close()

	hc := hdis.Conn{c}

	key := "object:1234567"
	hc.Set(key, "The value")
	value, _ := redis.String(hc.Get(key))

	fmt.Printf("%v", value)
	// Output: The value
}
