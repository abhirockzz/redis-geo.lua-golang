package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/go-redis/redis"
)

const redisgeoname string = "geotest"

var redisCoordinate string

func main() {

	flag.StringVar(&redisCoordinate, "redis-server", "localhost:6379", "redis server in the form host:port")
	flag.Parse()

	client := redis.NewClient(&redis.Options{Addr: redisCoordinate})
	fmt.Println("connected to redis at - " + redisCoordinate)

	client.GeoAdd(redisgeoname, &redis.GeoLocation{Name: "place1", Latitude: 1, Longitude: 2})
	client.GeoAdd(redisgeoname, &redis.GeoLocation{Name: "place2", Latitude: 3, Longitude: 4})
	client.GeoAdd(redisgeoname, &redis.GeoLocation{Name: "place3", Latitude: 5, Longitude: 6})
	fmt.Println("added locations to " + redisgeoname)

	bytes, err := ioutil.ReadFile("geo.lua")
	if err != nil {
		fmt.Println("error " + err.Error())
	}

	geoluascript := string(bytes)
	redisGeoLuaScript := redis.NewScript(geoluascript)

	result, execErr := redisGeoLuaScript.Eval(client, []string{redisgeoname},
		"GEOPATHLEN", "place1", "place2", "place3").Result()

	if execErr != nil {
		fmt.Println("error " + execErr.Error())
	}
	fmt.Println("GEOPATHLEN result - ", result)
}
