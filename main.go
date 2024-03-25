package main

import (
	"log"

	"google.golang.org/protobuf/proto"

	pb "zeus/msg"
)

func main() {
	s := &pb.Student{
		Name:   "Peng Jie",
		Age:    "24",
		Gender: "Male",
		Number: 99,
	}

	log.Println(
		s.GetName(),
		s.GetAge(),
		s.GetGender(),
		s.GetNumber(),
	)
	// 序列化
	data, err := proto.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(data)
	// 反序列化
	ss := &pb.Student{}
	err = proto.Unmarshal(data, ss)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(ss)
}
