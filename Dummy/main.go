package main

import (
	pb "Dummy/pb/go"
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"
	// pbj "Dummy/pb/java"
)

func getAccount() *pb.Account {
	return &pb.Account{
		Id:         1,
		Name:       "Sahil",
		Thumbnail:  []byte{},
		IsVarified: true,
		Height:     192.3,
		Phones:     []string{"9664904885", "7874956588"},
		EyeColor:   pb.EyeColor_EYE_COLOR_BLACK, //EyeColor: 1,
	}
}

func getDummy() *pb.Dummy {
	return &pb.Dummy{
		Id: 45,
	}
}

func getComplex() *pb.Complex {
	return &pb.Complex{
		Single: &pb.Simple{Id: 1, Name: "Sahil"},
		Multiple: []*pb.Simple{
			{Id: 2, Name: "Ram"},
			{Id: 3, Name: "Buddhu"},
		},
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"

	writeToFile(path, p)
	message := &pb.Account{}
	readFromFile(path, message)
}

func doToJSON(p proto.Message) string {

	jsonS := toJSON(p)
	fmt.Println(jsonS)
	return jsonS
}
func doFromJSON(jsonS string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonS, message)
	return message
}
func main() {
	fmt.Println(getDummy())
	fmt.Println(getAccount())
	fmt.Println(getComplex())
	doFile(getAccount())

	fmt.Println("***Handling JSON now...")
	jsonS := doToJSON(getAccount())
	message := doFromJSON(jsonS, reflect.TypeOf(pb.Account{}))
	fmt.Println(message)
}
