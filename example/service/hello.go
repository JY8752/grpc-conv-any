package service

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/JY8752/grpc-conv-any/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HelloService struct {
}

func (s *HelloService) Hello(ctx context.Context, empty *emptypb.Empty) (*pb.HelloResponse, error) {
	response := &pb.HelloResponse{
		Id:    1,
		Value: newAnyValue(),
	}
	return response, nil
}

func newAnyValue() *anypb.Any {
	res1 := &pb.Value1{StrValue: "value"}
	res2 := &pb.Value2{IntValue: 1111}
	res3 := &pb.Value3{WValue: &pb.WrapedValue{StrValue: "value"}}
	responses := [3]proto.Message{res1, res2, res3}

	//ランダムにレスポンスを決める
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(3)
	res := responses[index]

	any, err := anypb.New(res)
	if err != nil {
		log.Fatalf("failed new any type message err: %v", err)
	}
	return any
}
