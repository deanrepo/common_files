package main

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

const (
	followedNewActs = "followed_new_activities"
)

func main() {
	opts := []redis.DialOption{redis.DialDatabase(0), redis.DialPassword("123456")}
	c, err := redis.Dial("tcp", ":6379", opts...)
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}
	defer c.Close()

	ret, _ := c.Do("SET", "fleet", "truck1")
	fmt.Printf("%s\n", ret)

	ret, _ = c.Do("GET", "fleet")
	fmt.Printf("%s\n", ret)

	// test for nil return
	key := followedNewActs
	r, err := c.Do("HGET", key, "123")
	fmt.Printf("err: %v, result: %v", err, r)

}
