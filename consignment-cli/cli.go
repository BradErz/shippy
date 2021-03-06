// consignment-cli/cli.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	pb "github.com/BradErz/shippy/consignment-service/proto/consignment"
	"github.com/micro/go-micro/cmd"
	microclient "github.com/micro/go-micro/client"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &consignment)
	return consignment, nil
}

func main()  {
	cmd.Init()
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	// Contact the server and print the response
	consignment, err := parseFile(defaultFilename)

	if err != nil {
		log.Fatalf("failed to parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("failed to create consignment: %v", err)
	}

	log.Printf("Created: %t", r.Created)

	getAll, err :=  client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("failed to get consignments: %v", err)
	}

	for key, value := range getAll.Consignments {
		fmt.Printf("%d: %v", key, value)
	}
}