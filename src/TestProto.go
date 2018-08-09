package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"proto"
)

/**

 */
func main() {
	var lat int32 = 323
	var lon int32 = 3234
	a := &proto_dic.LocationData{
		Latitude:  &lat,
		Longitude: &lon,
	}
	data1, _ := proto.Marshal(a)
	fmt.Println(data1)
	data2 := &proto_dic.LocationData{}
	proto.Unmarshal(data1, data2)
}
