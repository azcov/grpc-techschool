package main

import (
	"context"
	"flag"
	"github.com/azcov/grpc-techschool/pb"
	"github.com/azcov/grpc-techschool/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server %s: %v", *serverAddress, err)
	}
	 laptopClient := pb.NewLaptopServiceClient(conn)
	 laptop :=  sample.NewLaptop()
	req := &pb.CreateLaptopRequest{Laptop: laptop}

	 // set timeout
	 ctx, cancel := context.WithTimeout(context.Background(),time.Second* 5)
	defer cancel()

	 res, err := laptopClient.CreateLaptop(ctx, req)
	 if err != nil {
	 	st, ok := status.FromError(err)
	 	if ok && st.Code() == codes.AlreadyExists {
	 		log.Fatal("laptop already exists")
		}
		 log.Fatal("cannot create laptop", err)
	 }

	 log.Printf("cerated laptop with id %s", res.Id)
}
