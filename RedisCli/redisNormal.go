package rediscli

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func RedisNormal(urlRedis string) {
	conn, err := redis.Dial("tcp", urlRedis)
	checkError(err)
	defer conn.Close()

	_, err = conn.Do(
		"HMSET",
		"key:1",
		"string1",
		"incentive_rules",
		"string2",
		"000998",
		"string3",
		"permen",
		"string4",
		"roti",
		"string5",
		1,
	)
	checkError(err)

	string1, err := redis.String(conn.Do("HGET", "key:1", "string1"))
	checkError(err)
	fmt.Print("Isi dari string1 : \n\n", string1)

	values, err := redis.StringMap(conn.Do("HGETALL", "key:1"))
	checkError(err)
	for k, v := range values {
		fmt.Println("Key:", k)
		fmt.Println("value:", v)
	}
}
