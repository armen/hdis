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

	fmt.Printf("%v\n", value)

	// Delete the hash
	_, err := hc.Do("HDEL", key)
	if err != nil {
		fmt.Println(err)
	}

	// Execute invalid hash command
	_, err = hc.Do("GET", "sample-key")
	if err == hdis.NotAHashCommandError {
		fmt.Println("Invalid hash command")
	} else {
		fmt.Println(err)
	}

	// Output:
	// The value
	// Invalid hash command
}
