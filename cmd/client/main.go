package main

import (
	"authorization/pkg/api"
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("not enought arguments")
	}

	login := flag.Arg(0)
	password := flag.Arg(1)

	connect, err := grpc.Dial(":1200", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := api.NewAuthorizationClient(connect)
	res, err := client.Register(context.Background(), &api.UserMeta{Login: login, Password: password})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetStatus())
}
