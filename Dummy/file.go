package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

func writeToFile(filename string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Error while serializing", err)
		return
	}

	if err = ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Fatalln("Error while writing file", err)
		return
	}

	fmt.Println("Written data in File")
}

func readFromFile(filename string, pb proto.Message) {
	in, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalln("Error while reading file", err)
		return
	}

	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Error while De-serializing", err)
		return
	}

	fmt.Println("Data in file :", pb)
}
