package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
	"proto"
)

func main() {
	locationData := &proto_dic.LocationData{}
	data := testRedis()
	proto.Unmarshal(data, locationData)
	jsonStr := ""
	PB2JSON(locationData, &jsonStr)
	fmt.Println(jsonStr)
}

func testRedis() []byte {
	c, _ := redis.Dial("tcp", "127.0.0.1:6379")
	t1, _ := redis.Bytes(c.Do("get", "LL_ALL:10019409551"))
	return t1
}

func PB2JSON(fromPb proto.Message, toJsonStr *string) error {

	// pb转json字符串
	jsonStr, err := json.Marshal(fromPb)
	if err == nil {
		fmt.Println(toJsonStr)
		fmt.Println(&toJsonStr)
		fmt.Println(*toJsonStr)
		*toJsonStr = string(jsonStr)
		fmt.Println(toJsonStr)
		fmt.Println(&toJsonStr)

	}

	return err
}
