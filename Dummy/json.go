package main

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func toJSON(pb proto.Message) string {
	option := protojson.MarshalOptions{
		Multiline: true,
	}

	// out, err := protojson.Marshal(pb)
	out, err := option.Marshal(pb)
	if err != nil {
		log.Fatalln("Error while converting to JSON", err)
		return ""
	}

	return string(out)
}

func fromJSON(json string, pb proto.Message) {
	if err := protojson.Unmarshal([]byte(json), pb); err != nil {
		log.Fatalln("Error while converting from JSON", err)
	}

	// fmt.Println("Data from JSON :", pb)
}
